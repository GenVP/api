package v8runner

import (
	"github.com/Khorevaa/go-v8runner/v8constants"
	"github.com/Khorevaa/go-v8runner/v8run"
	"github.com/Khorevaa/go-v8runner/v8tools"
	log "github.com/sirupsen/logrus"
)

// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type Error interface {
	Error() string
}

type КонфигураторИнтерфейс interface {
	ЗапускательКонфигуратора
	процедурыСозданияБазы
	процедурыЗагрузкиКонфигурации
	дополнительныеОбработчики
}

type ЗапускательКонфигуратора interface {
	ВыполнитьКомандуКонфигуратора() (err error)
	ВыполнитьКомандуСоздатьБазу() (err error)
	ВыполнитьКомандуПредприятие() (err error)
	ВыполнитьКоманду() (err error)

	УстановитьВерсиюПлатформы(строкаВерсияПлатформы string)
	КлючСоединенияСБазой() string
	УстановитьКлючСоединенияСБазой(КлючСоединенияСБазой string)
	УстановитьАвторизацию(Пользователь string, Пароль string)
	УстановитьПараметры(Параметры ...string)
	ДобавитьПараметры(Параметры ...string)
	ПолучитьВыводКоманды() (s string)
	ПроверитьВозможностьВыполнения() (err error)
}

type Конфигуратор struct {
	v8run.ЗапускательКонфигуратора
	временнаяБаза *ВременнаяБаза
}

// new func

func НовыйКонфигуратор() (conf *Конфигуратор) {

	conf = &Конфигуратор{}

	err := conf.УстановитьВерсиюПлатформы("8.3")

	if err != nil {
		log.Panicf("Не удалось установить версию платформы: %s", err)
	}

	conf.временнаяБаза = НоваяВременнаяБаза(v8tools.ВременныйКаталогСПрефисом(v8constants.TempDBname))
	conf.УстановитьКлючСоединенияСБазой(conf.КлючВременногоСоединенияСБазой())
	return conf
}

func (conf *Конфигуратор) КлючВременногоСоединенияСБазой() string {

	log.Debugf("Получение временного ключа соединения с базой: %s", conf.временнаяБаза.КлючСоединенияСБазой)

	return conf.временнаяБаза.КлючСоединенияСБазой
}

func (conf *Конфигуратор) ПроверитьВозможностьВыполнения() (err error) {

	log.Debugf("ВременнаяБазаСуществует: %v", conf.временнаяБаза.Cуществует)

	if conf.КлючСоединенияСБазой() == conf.КлючВременногоСоединенияСБазой() {

		if !conf.временнаяБаза.Cуществует {

			conf.временнаяБаза.ИнициализироватьВременнуюБазу()
		}

	}

	err = conf.ЗапускательКонфигуратора.ПроверитьВозможностьВыполнения()

	return

}

func (conf *Конфигуратор) ВыполнитьКомандуКонфигуратора() (err error) {

	err = conf.ПроверитьВозможностьВыполнения()

	if err != nil {
		return
	}

	err = conf.ЗапускательКонфигуратора.ВыполнитьКомандуКонфигуратора()

	return
}

func (conf *Конфигуратор) ВыполнитьКомандуСоздатьБазу() (err error) {

	err = conf.ПроверитьВозможностьВыполнения()

	if err != nil {
		return
	}

	err = conf.ЗапускательКонфигуратора.ВыполнитьКомандуСоздатьБазу()

	return
}

func (conf *Конфигуратор) ВыполнитьКомандуПредприятие() (err error) {

	err = conf.ПроверитьВозможностьВыполнения()

	if err != nil {
		return
	}

	err = conf.ЗапускательКонфигуратора.ВыполнитьКомандуПредприятие()

	return
}

func (conf *Конфигуратор) ВыполнитьКоманду() (err error) {

	err = conf.ПроверитьВозможностьВыполнения()

	if err != nil {
		return
	}

	err = conf.ЗапускательКонфигуратора.ВыполнитьКоманду()

	return
}
