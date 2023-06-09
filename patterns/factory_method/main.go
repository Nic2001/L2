package pattern

import "fmt"

//Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код
//Конструктор не может вернуть существующий экземпляр, а фабричный метод может

//Преимущества
//Выделяет код производства продуктов в	 одно место, упрощая поддержку кода.
//Упрощает добавление новых продуктов в программу.
//Реализует принцип открытости/закрытости.

//Недостатки
//Может привести к созданию больших параллельных иерархий

type IDeviceWrite interface {
	UseDevice()
}

type DeviceWrite struct {
	Name string
}

func (d *DeviceWrite) UseDevice() {
	fmt.Println("Use device: ", d.Name)
}

type Pen struct {
	DeviceWrite
	Ink int
}

func NewPen() IDeviceWrite {
	return &Pen{
		DeviceWrite: DeviceWrite{
			Name: "pen",
		},
		Ink: 10,
	}
}

type Pencil struct {
	DeviceWrite
	PecilLead int
}

func NewPencil() IDeviceWrite {
	return &Pencil{
		DeviceWrite: DeviceWrite{
			Name: "pencil",
		},
		PecilLead: 5,
	}
}
