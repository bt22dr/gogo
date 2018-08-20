package multiset

// 새로운 MultiSet을 생성하여 반환한다.
func NewMultiSet() map[string]int {
	return map[string]int{}
}

// Insert 함수는 집합에 val을 추가한다.
func Insert(m map[string]int, val string) {
	m[val]++
}

// Erase 함수는 집합에서 val을 제거한다. 집합에 val이 없는
// 경우에는 아무 일도 일어나지 않는다.
func Erase(m map[string]int, val string) {
	// 이 구현은 버그 있음. 아래 Erase 메소드 참고할것!
	if _, ok := m[val]; ok {
		m[val]--
	}
}

//Count 함수는 집합에 val이 들어 있는 횟수를 구한다.
func Count(m map[string]int, val string) int {
	return m[val]
}

// String 함수는 집합에 들어 있는 원소들을 { } 안에 빈 칸으로
// 구분하여 넣은 문자열을 반환한다.
func String(m map[string]int) string {
	s := "{"
	for k, v := range m {
		for i := 0; i < v; i++ {
			s += " " + string(k)
		}
	}
	s += " }"

	return s
}

type MultiSet map[string]int

func (m MultiSet) Insert(val string) {
	m[val]++
}

func (m MultiSet) Erase(val string) {
	if m[val] <= 1 {
		delete(m, val)
	} else {
		m[val]--
	}
}

func (m MultiSet) Count(val string) int {
	return m[val]
}

func (m MultiSet) String() string {
	s := "{ "
	for val, count := range m {
		for i := 0; i < count; i++ {
			s += val + " "
		}
	}
	return s + "}"
}
