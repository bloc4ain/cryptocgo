
   Markets                                                                                                                                                                                          Outputs
┏━━━━━━━━━━━┓                                                                                                                                                                                           ┌───────────┐
┃           ┃                                                                                                                                                                                           │           │
┃  Binance  ┣━━━┓                                                                                                                                                                                   ┌──▶│   CLI    │
┃           ┃   ┃                                                             Controller                   Processor                                                                                     Storage       │   │           │
┗━━━━━━━━━━━┛   ┃━━━━━━━━━━━                                             API          ┌─────────────────┐         ┌───────────────┐                                                                               ┌─────────────┐     ┌─────────────┐   │   └───────────┘
                ┃                                                      ◀────────────▶│ API calls       │         │  Transform    │                                                                            │             │     │             │   │
┏━━━━━━━━━━━┓   ┃                                                                     │                 │         │               │                                                                            │             │     │             │   │   ┌───────────┐
┃           ┃   ┃                                                        Event bus    │                 │ Data    │  Aggregate    │                                                                            │ * transform │     │  Rethink DB │   │   │           │
┃  KuCoin   ┣━━━╋━━━━━━━━━━━                                           ──────────────▶│ Event subs     │───────▶ │               │───────▶                                                                             ─▶│ * decorate │────▶│             │───┼──▶│   Web    │
┃           ┃   ┃                                                          │                 │         │               │                                                                       │ * filter    │     │             │   │   │           │
┗━━━━━━━━━━━┛   ┃                                             Control bus  │                 │         │               │                                                                       │             │     │             │   │   └───────────┘
                ┃                                           ◀──────────────│ Controll calls │         │               │                                                                         └─────────────┘     └─────────────┘   │
┏━━━━━━━━━━━┓   ┃━━━━━━━━━━━                                                          └─────────────────┘         └───────────────┘                                                                                                       │   ┌───────────┐
┃           ┃   ┃                                                                                                                                                                                                │   │           │
┃  Bittrex  ┣━━━┛                                                                                                                                                                                                └──▶│ Telegram │
┃           ┃                                                                                                                                                                                           │           │
┗━━━━━━━━━━━┛                                                                                                                                                                                           └───────────┘