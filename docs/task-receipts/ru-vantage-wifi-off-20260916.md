# Note3 Wi-Fi-off RU-vantage attempt — 2026-09-16

## Result

The Note3 was switched off Wi-Fi and stayed on its MTS mobile route, but the host-side
capture vantage could not be raised. No study rows or frame data were written.

## Evidence

- ADB serial: `4d00553d61ab90b7`, physical model `SM_N900`.
- Before the change, the phone reported `gsm.operator.alpha=MTS RUS` and
  `gsm.network.type=HSDPA`.
- `adb -s 4d00553d61ab90b7 shell su -c 'svc wifi disable'` returned 0 (with the known
  Samsung `resetreason` permission warning); afterward `settings get global wifi_on` was
  `0` and the phone route was `10.194.63.0/24 dev rmnet0 src 10.194.63.172`.
- A direct phone-side `ping -c 1 -W 3 1.1.1.1` succeeded at `417.957 ms`.
- `tools/adint-ru-vantage-keep --test` failed honestly: `wifi_on=0 rmnet0_inet=1`
  but `egress=unreachable`; its required `ruvantage` namespace did not exist.
- The documented RNDIS path was attempted. Host `enx0257356d3666` was moved into a
  temporary `ruvantage` namespace and configured as `192.168.42.130/24` with gateway
  `192.168.42.129`; the gateway did not answer and `curl https://ria.ru/` failed with
  DNS/no egress.
- On the phone, uid 0 could not assign `192.168.42.129/24` to `rndis0`:
  `RTNETLINK answers: Operation not permitted`. `iptables` NAT/filter setup was also
  denied by the live Android policy. This is the exact blocker, not a browser result.

## Cleanup

The temporary namespace and scoped host routes were removed. The host main route remained
`default via 192.168.8.1 dev enp42s0`. The phone was restored to `mtp,adb`; Wi-Fi remains
off (`wifi_on=0`) and the phone remains on `rmnet0`.

## Retry condition

Retry only after a host-visible RNDIS gateway is present or an operator-approved Android
tether/NAT path permits `rndis0` address and forwarding setup. Then rerun the real RU
vantage gate before any capture.
