package strconv

var __strconv_trash int

func Atoi(gs string) (int, error) {
	bts := []byte(gs)
	if len(bts) == 0 {
		return 0,nil
	}
	var b byte
	var n int
	var i int
	var isMinus bool
	for i, b = range bts {
		if b == '.' {
			return 0,nil // @FIXME all no number should return error
		}
		if b == '-' {
			isMinus = true
			continue
		}
		var x byte = b - byte('0')
		n  = n * 10
		n = n + int(x)
	}
	if isMinus {
		n = -n
	}
	__strconv_trash = i
	return n, nil
}

func Itoa(i int) string {
	var r []byte
	var tmp []byte
	var isMinus bool

	// open(2) returs  0xffffffff 4294967295 on error.
	// I don't understand this yet.
	if i > 2147483648 {
		i = i - 2147483648*2
	}


	if i < 0 {
		i = i * -1
		isMinus = true
	}
	for i>0 {
		mod := i % 10
		tmp = append(tmp, byte('0') + byte(mod))
		i = i /10
	}

	if isMinus {
		r = append(r, '-')
	}

	for j:=len(tmp)-1;j>=0;j--{
		r = append(r, tmp[j])
	}

	if len(r) == 0 {
		return "0"
	}
	return string(r)
}
