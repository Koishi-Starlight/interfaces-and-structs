package electronics

type Phone interface {
	Brand() string
	Model() string
	Type() string
}
type StationPhone interface {
	ButtonsCount() int
}
type Smartphone interface {
	OS() string
}

// Apple Phone
type applePhone struct {
	model string
}

func (a applePhone) Brand() string {
	return "Apple"
}
func (a applePhone) Model() string {
	return a.model
}
func (a applePhone) Type() string {
	return "Smartphone"
}
func (a applePhone) OS() string {
	return "iOS"
}
func NewApplePhone(model string) applePhone {
	return applePhone{model: model}
}

// Android Phone
type androidPhone struct {
	brand string
	model string
}

func (a androidPhone) Brand() string {
	return a.brand
}
func (a androidPhone) Model() string {
	return a.model
}
func (a androidPhone) Type() string {
	return "Smartphone"
}
func (a androidPhone) OS() string {
	return "Android"
}
func NewAndroidPhone(brand, model string) androidPhone {
	return androidPhone{brand: brand, model: model}
}

// Radio Phone
type radioPhone struct {
	brand   string
	model   string
	buttons int
}

func (r radioPhone) Brand() string {
	return r.brand
}
func (r radioPhone) Model() string {
	return r.model
}
func (r radioPhone) Type() string {
	return "Station"
}
func (r radioPhone) ButtonsCount() int {
	return r.buttons
}
func NewRadioPhone(brand, model string, buttons int) radioPhone {
	return radioPhone{brand: brand, model: model, buttons: buttons}
}
