package safe_validator

type t0 struct {
	I0   int               `valid:"not_zero;range:-10~10;enum:1,2,3"`
	F0   float64           `valid:"not_zero;range:-0.1~0.1,enum:-0.01,0.01,0.09"`
	S0   string            `valid:"not_zero;len:1~10;rule:number"`
	IPtr *int              `valid:"not_nil;range:~2;enum:3,2,1"`
	SPtr *string           `valid:"not_nil;len:10~20;enum:a,c,b"`
	Arr0 []int             `valid:"not_zero;len:10"`
	Map0 map[string]string `valid:"not_zero;len:20"`
}
