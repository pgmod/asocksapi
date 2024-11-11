### Код, который создает один shared прокси с корпоративным сервером и ротацией после каждого запроса в случайной стране

``` golang

api := asocksapi.NewApi(
  os.Getenv("ASOCKS_API_KEY"),
)

prox_, _ := api.CreatePorts(
  asocksapi.CreatePortsReq{
    CountryCode:      asocksapi.Countries[rand.Intn(len(asocksapi.Countries))],
    TypeId:           asocksapi.ROTATE_CONNECTION,
    ProxyTypeId:      asocksapi.CORPORATE,
    ServerPortTypeId: asocksapi.SHARED,
    Count:            1,
    Name:             "asocs-auto",
  },
)

if len(prox_.Data) == 0 {
  bot.Logger.Error("Не удалось создать прокси")
  os.Exit(10)
} else {
  Logger.Info("Прокси создан: ", prox_.Data[0].Country)
}

```
