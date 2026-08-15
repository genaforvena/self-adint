package net.selfadint.probe;

import android.os.Bundle;
import android.view.Gravity;
import android.widget.ScrollView;
import android.widget.TextView;

import androidx.appcompat.app.AppCompatActivity;

import org.prebid.mobile.BannerAdUnit;
import org.prebid.mobile.Host;
import org.prebid.mobile.PrebidMobile;
import org.prebid.mobile.ResultCode;
import org.prebid.mobile.api.data.BidInfo;
import org.prebid.mobile.api.data.InitializationStatus;
import org.prebid.mobile.rendering.listeners.SdkInitializationListener;

/**
 * The whole app. It exists to make the SDK describe this phone to our own Prebid Server,
 * and to do nothing else.
 *
 * The point of the step is that WE do not write the bid request. Every field in it —
 * device.ifa, the make and model, the OS version, the screen, the connection type, the
 * locale, the carrier — is filled in by the Prebid Mobile SDK reading the device it is
 * running on. That is why this file contains no JSON: the moment we hand-assemble a
 * request, the artifact stops being evidence about the phone and becomes a document about
 * our own imagination.
 *
 * There is no UI beyond a status line and nothing is stored on the device.
 */
public class MainActivity extends AppCompatActivity {

    private TextView out;

    private void say(String s) {
        runOnUiThread(() -> out.append(s + "\n\n"));
    }

    @Override
    protected void onCreate(Bundle saved) {
        super.onCreate(saved);

        out = new TextView(this);
        out.setTextSize(13f);
        out.setPadding(24, 48, 24, 24);
        out.setGravity(Gravity.START);
        ScrollView sv = new ScrollView(this);
        sv.addView(out);
        setContentView(sv);

        say("adint probe\ntarget: " + BuildConfig.PBS_HOST);

        // Our own server, on the LAN. Nothing here talks to any exchange.
        PrebidMobile.setPrebidServerAccountId("adint-local");
        PrebidMobile.setPrebidServerHost(Host.createCustomHost(BuildConfig.PBS_HOST));
        PrebidMobile.setShareGeoLocation(false);   // geo is not what this step is measuring
        PrebidMobile.setTimeoutMillis(3000);

        PrebidMobile.initializeSdk(getApplicationContext(), new SdkInitializationListener() {
            @Override
            public void onInitializationComplete(InitializationStatus status) {
                say("sdk init: " + status);
                fetch();
            }
        });
    }

    private void fetch() {
        BannerAdUnit unit = new BannerAdUnit("adint-imp-1", 300, 250);
        say("requesting demand — the SDK is now composing the bid request");
        // fetchDemand(OnFetchDemandResult) — BidInfo carries the ResultCode plus the
        // targeting keywords a real publisher would forward to its ad server.
        unit.fetchDemand((BidInfo info) -> {
            ResultCode resultCode = info.getResultCode();
            say("result: " + resultCode
                    + (resultCode == ResultCode.NO_BIDS
                       ? "\n(NO_BIDS is the EXPECTED outcome: the local server's only bidder"
                         + " is a stub that answers 204. The artifact is the request, not a bid.)"
                       : ""));
            say("Now read the capture on the server:\n"
                    + "  ls ~/self-adint/data/bidreq/\n"
                    + "The .meta.json beside it carries this phone's LAN address — a capture whose"
                    + " peer is 127.0.0.1 was made on the server and is not this device.");
        });
    }
}
