package service

//等待实现的接口
type Service interface {
	Add (a,b int)int
}

type ArithmeticService struct {}

func (a *ArithmeticService) Add(n1,n2 int)int  {
	return n1+n2
}
