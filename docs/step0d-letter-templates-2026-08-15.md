# Шаблоны писем субъекта данных — три опоры, три текста

2026-08-15. Классификация адресатов: `step0d-recipients-eu-entity-2026-08-15-ru.md`.
Фактура на каждого: `tools/adint-dsar-evidence` → `data/dsar-evidence.tsv`.
Отправка: `tools/adint-mail-send` (адресату и копия оператору — одно действие).

**Язык писем — английский.** Все оставшиеся адресаты иностранные; русские исключены решением
оператора. Русский здесь только в комментариях для оператора.

**Один шаблон на всех — это ровно та ошибка, из-за которой запрос снимают.** Класс A опирается на
закон и называет европейское юрлицо. Класс B не может этого утверждать — там формулировка
**условная**. Класс C не имеет законной опоры вообще, и это говорится прямо: опора — их собственное
опубликованное обещание.

---

## ВХОД, КОТОРОГО ЕЩЁ НЕТ: рекламный идентификатор

Ни в одном снимке GAID не виден — снималка ловит хосты, время и байты, но не содержимое TLS.
Без идентификатора запрос беспредметен: адресат физически не сможет найти запись, даже если
захочет.

Взять его с телефона, любым из двух:
- **Настройки → Google → Реклама → «Ваш рекламный идентификатор»** (или Настройки → Конфиденциальность → Реклама);
- по adb: `adb shell settings get secure advertising_id` (на части прошивок пусто — тогда первый способ).

**Это тот самый идентификатор, который письмо свяжет с почтой** — поэтому ящик отдельный и
нейтральный. У Bidease присылка GAID вообще их **обязательное** условие обработки запроса.

---

## Общая рамка всех трёх писем

1. **Тема** — короткая, чтобы попала в правильную очередь, и без имени в ней.
2. **Кто пишет** — имя и фамилия заявителя в теле (не в адресе).
3. **Идентификатор** — GAID, и прямо сказано, что это идентификатор, по которому у них ведётся запись.
4. **Доказательный блок** — хосты, время первого и последнего контакта, байты **отправленные**
   с устройства, приложение-источник. Каждая строка из его снимка, **ни одного выдуманного факта**.
5. **Запрос** — не «есть ли что-нибудь», а перечень: сегменты и их **названия**, источники,
   получатели, срок хранения, связанные идентификаторы.
6. **Удаление** — в том же письме, и **с требованием перечислить**, что удержано и на каком
   основании. Иначе придёт «часть данных сохранена» без предмета.
7. **Срок ответа** — назван, чтобы молчание стало измеримым, а не бессрочным ожиданием.

---

## ШАБЛОН A — названное европейское юрлицо (опора: GDPR Art 3(1))

Адресаты: **Xiaomi** (Xiaomi Technology Netherlands B.V. — совместный контролёр для EEA, самый
прочный случай в списке); **Google** — с оговоркой, см. ниже.

> **Subject:** Data subject access and erasure request — advertising identifier `<GAID>`
>
> To the Data Protection Officer,
>
> My name is `<ИМЯ ФАМИЛИЯ>`. I am making a request concerning personal data your company processes
> about my mobile device, identified by the Google Advertising ID **`<GAID>`**.
>
> **Applicable law.** I am aware that I am not located in the European Union. I am not relying on
> Article 3(2) of Regulation (EU) 2016/679. I am relying on **Article 3(1)**: the Regulation applies
> to the processing of personal data "in the context of the activities of an establishment of a
> controller or processor in the Union, regardless of whether the processing takes place in the
> Union or not". Your own privacy policy names **`<ЕВРОПЕЙСКОЕ ЮРЛИЦО, адрес>`** as
> `<a joint controller for the EEA / an entity providing services in the EU>`. This request concerns
> processing carried out in the context of that establishment's activities.
>
> **Evidence that processing took place.** The following exchanges between my device and hosts
> operated by your company were recorded on my own device on 15 August 2026 (local time UTC+3):
>
> ```
> <ДОКАЗАТЕЛЬНЫЙ БЛОК из dsar-evidence.tsv: хост · соединений · байт отправлено · байт получено ·
>  первый контакт · последний контакт · приложение-источник>
> ```
>
> **Under Article 15 I request:**
> 1. confirmation as to whether personal data concerning me is being processed, and if so, a copy of it;
> 2. **the audience segments, interest categories or taxonomy entries associated with this identifier — by their names, not by count**;
> 3. the purposes of the processing;
> 4. the categories of personal data concerned;
> 5. the recipients or categories of recipient to whom the data has been or will be disclosed, including any recipients in third countries;
> 6. the envisaged retention period, or the criteria used to determine it;
> 7. all available information as to the **source** of the data where it was not collected from me directly;
> 8. any other identifiers your systems have linked to this advertising identifier.
>
> **Under Article 17 I also request erasure** of the personal data associated with this identifier.
>
> I am aware that erasure may lawfully be refused in whole or in part where you are required to
> retain data under Union or Member State law. **If you refuse erasure on that basis, please state
> which categories of data are retained, for how long, and the specific legal obligation relied
> upon.** A statement that "some data has been retained", without identifying what, does not allow
> me to assess the response.
>
> Please reply to this address. I understand the response is due within one month of receipt under
> Article 12(3), and that this may be extended by two further months where necessary, provided I am
> informed of the extension and the reasons for it within the first month.
>
> `<ИМЯ ФАМИЛИЯ>`

**Оговорка по Google, и её надо держать при отправке.** Google Ireland Limited назван в их политике
как компания, оказывающая услуги в ЕС, но эта же политика относит пользователей **вне EEA** к
Google LLC. То есть опора Art 3(1) доступна, но **спорна**, и ответ «вашим контролёром является
Google LLC, GDPR к вам не применяется» — вероятный и законный исход. Он тоже измерение: он называет
контролёра. Письмо Google уходит по шаблону A, но с ожиданием этого ответа, а не как к Xiaomi.

---

## ШАБЛОН B — только представитель по Art 27 (опора: условная + их политика)

Адресаты: **Liftoff** (Liftoff GmBH, Berlin — назван Representative), **AppsFlyer** (AppsFlyer
Germany GmbH, Berlin — назван EU representative).

Отличие от A ровно одно и оно принципиальное: **утверждать наличие establishment мы не можем**.
Представитель по Art 27 назначается именно потому, что establishment'а нет. Поэтому опора
формулируется **условно**, и рядом сразу ставится вторая, независимая от неё.

> **Subject:** Data subject access and erasure request — advertising identifier `<GAID>`
>
> To the Data Protection Officer,
>
> My name is `<ИМЯ ФАМИЛИЯ>`. I am making a request concerning personal data your company processes
> about my mobile device, identified by the Google Advertising ID **`<GAID>`**.
>
> **Basis for this request.** I am not located in the European Union and I am not relying on
> Article 3(2) of Regulation (EU) 2016/679. Your privacy policy names
> **`<ЮРЛИЦО, адрес>`** as your representative in the EEA.
>
> **To the extent** that any of the processing described below is carried out in the context of the
> activities of an establishment of your company in the Union, I make this request under
> **Articles 15 and 17** of that Regulation. **In any event**, I make the same request under your
> own published privacy policy, which sets out the rights you undertake to provide to individuals
> whose data you process, and I ask you to honour it on its own terms.
>
> If you consider that neither basis applies, **please say so explicitly and state which law or
> policy you consider to govern the data you hold about this identifier.**
>
> **Evidence that processing took place.**
>
> ```
> <ДОКАЗАТЕЛЬНЫЙ БЛОК>
> ```
>
> **I request:**
> `<тот же перечень из восьми пунктов, что в шаблоне A>`
>
> **I also request erasure** of the personal data associated with this identifier. If you refuse
> erasure in whole or in part, please state which categories are retained, for how long, and on
> what basis.
>
> Please reply to this address within one month of receipt.
>
> `<ИМЯ ФАМИЛИЯ>`

**Отдельно по AppsFlyer.** По их собственному тексту в отношении данных конечных пользователей они
**процессор**, а не контролёр, и прямо направляют такие запросы к своим клиентам-приложениям. Это
законный и ожидаемый ответ. Он всё равно полезен: чтобы переадресовать, им придётся **назвать
клиента** — то есть приложение, от имени которого они его обрабатывают. Доказательный блок уже
называет приложения (`com.sofascore.results`, `ru.hh.android`, `ru.yandex.yandexmaps`), и в письме
стоит попросить подтвердить именно это соответствие.

---

## ШАБЛОН C — европейского присутствия нет (опора: их собственное обещание, не закон)

Адресаты: **Pangle/ByteDance** (Bytedance Pte. Ltd., Сингапур; европейское юрлицо не названо
вообще), **Moloco** (Moloco, Inc.), **Bidease** (Bidease Inc.).

Здесь **законного рычага нет**, и письмо не притворяется, что он есть — притворство даёт повод
отклонить всё целиком. Опора называется своим именем: **их опубликованное обещание**. У Bidease оно
записано их же словами, и цитата идёт в письмо.

> **Subject:** Data subject access and erasure request — advertising identifier `<GAID>`
>
> To the Privacy Team,
>
> My name is `<ИМЯ ФАМИЛИЯ>`. I am making a request concerning personal data your company processes
> about my mobile device, identified by the Google Advertising ID **`<GAID>`**.
>
> **Basis for this request.** I am located outside the European Union and I make no claim under
> Regulation (EU) 2016/679. I am relying on **your own published privacy policy**, in which you
> undertake to provide access and deletion rights to individuals whose data you process.
> `<ДЛЯ BIDEASE: Your policy states that Bidease "voluntarily commits to comply with The EU General
> Data Protection Regulation (GDPR)". I am asking you to honour that commitment as written.>`
>
> **Evidence that processing took place.**
>
> ```
> <ДОКАЗАТЕЛЬНЫЙ БЛОК>
> ```
>
> **I request:**
> `<тот же перечень из восьми пунктов>`
>
> **I also request deletion** of the personal data associated with this identifier. If you are
> required to retain any part of it, please state which categories, for how long, and on what basis.
>
> Please confirm receipt and reply to this address within 30 days.
>
> `<ИМЯ ФАМИЛИЯ>`

**Для Bidease — их процедура, а не наша.** Их политика прямо требует: *«To opt-out please email us
at privacy@bidease.com. In order for us to proceed with the request, your email must contain: —
Your advertising ID (e.g. IDFA / GAID / etc.)»*. То есть присылка GAID — их обязательное условие,
и адрес получателя фиксирован: `privacy@bidease.com`, DPO Boris Abaev.

---

## Адреса и что ещё нужно перед отправкой

| адресат | класс | адрес из их же документов |
|---|---|---|
| Xiaomi | A | `https://privacy.mi.com/support` + почтовый адрес Xiaomi Technology Netherlands B.V., The Hague |
| Google | A/спорно | форма Google по запросам субъектов |
| Liftoff | B | `privacy@liftoff.io` / DPO-адрес из политики; EEA rep: Liftoff GmBH, Berlin |
| AppsFlyer | B | `privacy@appsflyer.com`; EU rep: AppsFlyer Germany GmbH, Berlin |
| Pangle/ByteDance | C | `europe_privacy@pangleglobal.com` (их же региональный адрес) |
| Moloco | C | форма `moloco.com/contact-us` |
| Bidease | C | `privacy@bidease.com` (DPO Boris Abaev) |
| **Adjust** | **не установлен** | **adjust.com отдаёт 429 с этой ноды — классифицировать до письма, не после** |
| netpygrid | D | адресата нет |

Открытые входы, без которых письма не уходят:

1. **GAID** — с телефона (см. выше). Без него письмо беспредметно.
2. **Имя и фамилия заявителя** в том виде, в каком он хочет их указать.
3. **Адрес ящика** — заводит оператор, пароль приложения ложится в `~/.mesh/adint-mail.env`.
4. **Adjust** — перепроверить наличие европейского юрлица.

Точные адреса из таблицы перед отправкой **сверяются с текущей страницей политики каждого
адресата**, а не берутся отсюда: этот файл — снимок на 2026-08-15, и адрес для запросов субъекта
меняется чаще, чем сама политика.
