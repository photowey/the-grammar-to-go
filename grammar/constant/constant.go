//constant.go

package constant

const (
    Monday = iota + 1
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
    Sunday
)

const (
    Open = 1 << iota
    Close
    Pending
)
