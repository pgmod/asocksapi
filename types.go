package asocksapi

type ProxyTypeId int

const (
	// Резидентские
	//
	// домашние пк
	RESEDENTIAL ProxyTypeId = 1
	// Все типы прокси
	ALL ProxyTypeId = 2
	// Мобильные
	//
	// сотовая связь
	MOBILE ProxyTypeId = 3
	// Корпоративные
	//
	// сервера организаций
	CORPORATE ProxyTypeId = 4
)

type TypeId int

const (
	// Разрывать соединение
	//
	// Если IP-адрес отключится, соединение будет потеряно до тех пор, пока IP-адрес не восстановится
	KEEP_PROXY TypeId = 1
	// Сохранять соединение
	//
	// Мы подберем другой ближайший прокси из той же маски подсети или тот-же ASN. Для конечного сайта будет выглядеть, что у вас сменился IP, но вы остались у того же провайдера и это не выглядит подозрительным
	KEEP_CONNECTION TypeId = 2
	// Ротировать
	//
	// Смена IP-адреса по запросу и при переподключении
	ROTATE_CONNECTION TypeId = 3
)

type ServerPortTypeId int

const (
	//СТОИМОСТЬ ТАКОГО ПОРТА СОСТАВЛЯЕТ $0.01 В ДЕНЬ
	//
	//Авторизация по белому списку или авторизация по логину и паролю (Все, у кого есть пароль для входа или его IP, сохраняются в белом списке)
	DEDICATED ServerPortTypeId = 1
	// Только авторизация по логину и паролю
	//
	// Прокси будет работать только по логину и паролю, белый список IP будет игнорироваться
	SHARED ServerPortTypeId = 0
)
