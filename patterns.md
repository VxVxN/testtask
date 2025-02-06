### Паттерн «Фасад» (Facade)

**Применение:** Паттерн «Фасад» предоставляет упрощенный интерфейс к сложной подлежащей системе, позволяя клиентам взаимодействовать с ней с минимальной сложностью. Хорошим примером может служить система, которая имеет сложные подсистемы для работы с данными, отображения пользовательского интерфейса и взаимодействия с сетью. Вместо того чтобы взаимодействовать с каждой подсистемой по отдельности, клиентские приложения могут использовать фасад, который предоставляет единый интерфейс. Это не только упрощает использование системы, но и позволяет скрыть сложную логику и детали реализации от клиентов, обеспечивая более чистый и понятный код.

**Плюсы:**
- Упрощение использования системы.
- Снижение зависимости кода от сложных подсистем.

**Минусы:**
- Может скрыть сложность системы, что усложнит её изменение в будущем.

**Пример на Go:**
```go
package main

import "fmt"

type SubsystemA struct{}

func (a *SubsystemA) OperationA() {
    fmt.Println("SubsystemA: Operation A")
}

type SubsystemB struct{}

func (b *SubsystemB) OperationB() {
    fmt.Println("SubsystemB: Operation B")
}

type Facade struct {
    subsystemA *SubsystemA
    subsystemB *SubsystemB
}

func NewFacade() *Facade {
    return &Facade{
        subsystemA: &SubsystemA{},
        subsystemB: &SubsystemB{},
    }
}

func (f *Facade) Operation() {
    f.subsystemA.OperationA()
    f.subsystemB.OperationB()
}

func main() {
    facade := NewFacade()
    facade.Operation()
}
```

### Паттерн «Строитель» (Builder)

**Применение:** Паттерн «Строитель» позволяет создавать сложные объекты пошагово. Вместо того чтобы передавать все параметры в конструктор, что может привести к неразборчивым вызовам, строители могут создавать объект поэтапно, устанавливая лишь те параметры, которые необходимы. к примеру, при создании экземпляра человека, у вас может быть множество параметров, таких как имя, возраст, адрес и т.д., которые задаются не одновременно. Строитель позволяет изначально создать базовый объект, а затем последовательно добавлять все необходимые атрибуты. Это упрощает создание объектов, которые имеют много параметров, и улучшает читаемость кода.

**Плюсы:**
- Упрощает создание сложных объектов.
- Позволяет использовать один и тот же код для создания различных объектов.

**Минусы:**
- Увеличивает количество классов.

**Пример на Go:**
```go
package main

import "fmt"

type Product struct {
    partA string
    partB string
}

type Builder interface {
    BuildPartA()
    BuildPartB()
    GetProduct() Product
}

type ConcreteBuilder struct {
    product Product
}

func (b *ConcreteBuilder) BuildPartA() {
    b.product.partA = "Part A"
}

func (b *ConcreteBuilder) BuildPartB() {
    b.product.partB = "Part B"
}

func (b *ConcreteBuilder) GetProduct() Product {
    return b.product
}

type Director struct {
    builder Builder
}

func (d *Director) Construct() {
    d.builder.BuildPartA()
    d.builder.BuildPartB()
}

func main() {
    builder := &ConcreteBuilder{}
    director := Director{builder: builder}

    director.Construct()
    product := builder.GetProduct()
    fmt.Printf("Product built with: %s, %s\n", product.partA, product.partB)
}
```

### Паттерн «Посетитель» (Visitor)

**Применение:** Паттерн «Посетитель» позволяет добавлять новые операции на элементы объекта, не изменяя самих объектов. Это особенно полезно, когда существуют сложные иерархии классов, и вы хотите применить различные операции к их элементам. Например, представьте себе структуру классов, представляющих различные узлы дерева. Используя паттерн «Посетитель», можно создать класс посетителя, который будет реализовывать различные действия, такие как агрегация, анализ или изменение данных, применяемые ко всем узлам дерева. В дополнение к этому, это позволяет добавлять новые операции, не вмешиваясь в существующий код узлов.

**Плюсы:**
- Упрощает добавление новых операций.
- Изолирует операции от самих объектов.

**Минусы:**
- Усложняет структуру кода.
- Требует изменения кода, если добавляются новые элементы.

**Пример на Go:**
```go
package main

import "fmt"

type Element interface {
    Accept(visitor Visitor)
}

type Visitor interface {
    VisitConcreteElementA(elementA *ConcreteElementA)
    VisitConcreteElementB(elementB *ConcreteElementB)
}

type ConcreteElementA struct{}

func (e *ConcreteElementA) Accept(visitor Visitor) {
    visitor.VisitConcreteElementA(e)
}

type ConcreteElementB struct{}

func (e *ConcreteElementB) Accept(visitor Visitor) {
    visitor.VisitConcreteElementB(e)
}

type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(elementA *ConcreteElementA) {
    fmt.Println("Visiting ConcreteElementA")
}

func (v *ConcreteVisitor) VisitConcreteElementB(elementB *ConcreteElementB) {
    fmt.Println("Visiting ConcreteElementB")
}

func main() {
    elements := []Element{&ConcreteElementA{}, &ConcreteElementB{}}
    visitor := &ConcreteVisitor{}

    for _, element := range elements {
        element.Accept(visitor)
    }
}
```

### Паттерн «Комманда» (Command)

**Применение:** Паттерн «Команда» инкапсулирует запрос как объект, позволяя параметризовать клиентов с очередями, журналами и поддерживать операции, которые можно отменить. Например, в текстовом редакторе можно реализовать команды для редактирования текста (копирование, вставка, удаление). Каждая команда будет представлена как отдельный объект, который сможет выполнять свои действия по нажатию кнопок, а также поддерживать возможность отмены последних изменений.

**Плюсы:**
- Упрощает добавление новых команд.
- Позволяет реализовать операции отмены/повтора.

**Минусы:**
- Увеличивает количество классов.

**Пример на Go:**
```go
package main

import "fmt"

type Command interface {
    Execute()
}

type Light struct{}

func (l *Light) TurnOn() {
    fmt.Println("Light is ON")
}

func (l *Light) TurnOff() {
    fmt.Println("Light is OFF")
}

type TurnOnCommand struct {
    light *Light
}

func (c *TurnOnCommand) Execute() {
    c.light.TurnOn()
}

type TurnOffCommand struct {
    light *Light
}

func (c *TurnOffCommand) Execute() {
    c.light.TurnOff()
}

type RemoteControl struct {
    command Command
}

func (r *RemoteControl) PressButton() {
    r.command.Execute()
}

func main() {
    light := &Light{}
    turnOn := &TurnOnCommand{light: light}
    turnOff := &TurnOffCommand{light: light}

    remote := RemoteControl{command: turnOn}
    remote.PressButton()

    remote.command = turnOff
    remote.PressButton()
}
```

### Паттерн «Цепочка вызовов» (Chain of Responsibility)

**Применение:** Паттерн «Цепочка вызовов» помогает избежать жесткой привязки отправителя запроса к его получателю, позволяя передавать запрос вдоль цепочки обработчиков. Каждый обработчик в цепочке может либо обработать запрос, либо передать его следующему обработчику. Это особенно полезно в системах, где необходимо обрабатывать различные типы запросов, такие как события пользовательского интерфейса или сообщения в системе. Например, в приложении, обрабатывающем различные типы событий (клик мышью, нажатие клавиш), каждый обработчик может проверять, подходит ли ему этот тип события, и, если нет, передавать его дальше.

**Плюсы:**
- Упрощает добавление новых обработчиков.
- Снижает зависимость между отправителями и получателями запросов.

**Минусы:**
- Может привести к трудно отладимым цепочкам.

**Пример на Go:**
```go
package main

import "fmt"

type Handler interface {
    SetNext(handler Handler) Handler
    Handle(request string) string
}

type BaseHandler struct {
    next Handler
}

func (h *BaseHandler) SetNext(handler Handler) Handler {
    h.next = handler
    return handler
}

func (h *BaseHandler) Handle(request string) string {
    if h.next != nil {
        return h.next.Handle(request)
    }
    return ""
}

type ConcreteHandlerA struct {
    BaseHandler
}

func (h *ConcreteHandlerA) Handle(request string) string {
    if request == "A" {
        return "HandlerA: handling request A"
    }
    return h.BaseHandler.Handle(request)
}

type ConcreteHandlerB struct {
    BaseHandler
}

func (h *ConcreteHandlerB) Handle(request string) string {
    if request == "B" {
        return "HandlerB: handling request B"
    }
    return h.BaseHandler.Handle(request)
}

func main() {
    handlerA := &ConcreteHandlerA{}
    handlerB := &ConcreteHandlerB{}

    handlerA.SetNext(handlerB)

    fmt.Println(handlerA.Handle("A"))
    fmt.Println(handlerA.Handle("B"))
    fmt.Println(handlerA.Handle("C"))
}
```

### Паттерн «Фабричный метод» (Factory Method)

**Применение:** Паттерн «Шаблонный метод» определяет скелет алгоритма в методе, откладывая определение некоторых шагов на подклассы. Поскольку алгоритм определен в базовом классе, подклассы могут переопределять только те шаги, которые им нужны. Например, в игре можно иметь абстрактный класс «уровень», который определяет структуру уровня (создание врагов, сбор предметов), но конкретные детали реализации - создаются в подклассах для каждого специфического уровня.

**Плюсы:**
- Упрощает процесс создания объектов.
- Способствует инкапсуляции создания объектов.

**Минусы:**
- Увеличивает количество классов.

**Пример на Go:**
```go
package main

import "fmt"

type Product interface {
    Use() string
}

type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
    return "Using Product A"
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
    return "Using Product B"
}

type Creator interface {
    FactoryMethod() Product
}

type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) FactoryMethod() Product {
    return &ConcreteProductA{}
}

type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) FactoryMethod() Product {
    return &ConcreteProductB{}
}

func main() {
    creators := []Creator{&ConcreteCreatorA{}, &ConcreteCreatorB{}}

    for _, creator := range creators {
        product := creator.FactoryMethod()
        fmt.Println(product.Use())
    }
}
```

### Паттерн «Стратегия» (Strategy)

**Применение:** Паттерн «Стратегия» используется для определения семейства алгоритмов и их замены без изменения кода, который их использует. Например, в приложении для обработки платежей можно использовать разные стратегии для обработки различных типов платежей: кредитные карты, PayPal, криптовалюты. Каждый алгоритм будет реализован в своём классе, а контекст будет использовать нужную стратегию в зависимости от выбранного метода оплаты.

**Плюсы:**
- Упрощает добавление новых стратегий.
- Снижает зависимость между классами.

**Минусы:**
- Увеличивает количество классов.

**Пример на Go:**
```go
package main

import "fmt"

type Strategy interface {
    DoOperation(int, int) int
}

type Addition struct{}

func (a *Addition) DoOperation(x, y int) int {
    return x + y
}

type Subtraction struct{}

func (s *Subtraction) DoOperation(x, y int) int {
    return x - y
}

type Context struct {
    strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
    c.strategy = strategy
}

func (c *Context) ExecuteStrategy(x, y int) int {
    return c.strategy.DoOperation(x, y)
}

func main() {
    context := &Context{}

    context.SetStrategy(&Addition{})
    fmt.Println("Addition:", context.ExecuteStrategy(5, 3))

    context.SetStrategy(&Subtraction{})
    fmt.Println("Subtraction:", context.ExecuteStrategy(5, 3))
}
```

### Паттерн «Состояние» (State)

**Применение:** Паттерн «Состояние» позволяет объекту изменять свое поведение в зависимости от его внутреннего состояния. Один из примеров - автомат с напитками, который проявляет разное поведение в зависимости от состояния (например, «ожидание монет», «выдача напитка», «недостаточно средств»). Каждое состояние реализуется в своем классе, и сам автомат просто делегирует вызовы соответствующему состоянию.

**Плюсы:**
- Упрощает код, убирая множество условных операторов.
- Улучшает расширяемость системы.

**Минусы:**
- Увеличивает количество классов.

**Пример на Go:**
```go
package main

import "fmt"

type State interface {
    DoAction(context *Context)
}

type Context struct {
    state State
}

func (c *Context) SetState(state State) {
    c.state = state
}

func (c *Context) Request() {
    c.state.DoAction(c)
}

type StartState struct{}

func (s *StartState) DoAction(context *Context) {
    fmt.Println("Player is in Start State")
    context.SetState(&StopState{})
}

type StopState struct{}

func (s *StopState) DoAction(context *Context) {
    fmt.Println("Player is in Stop State")
    context.SetState(&StartState{})
}

func main() {
    context := &Context{state: &StartState{}}
    context.Request()
    context.Request()
}
```