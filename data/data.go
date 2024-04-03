package data

type Square struct {
  Num int
  IsBlack bool
  IsActive bool
}

type Data struct {
  Squares []Square
  ActiveSquares []*Square
}

func CreateData() *Data {
  var squares []Square

  for i := 0; i < 8; i++ {
    for j := 0; j < 8; j++ {
      isBlack := true

      if i % 2 == 0 && j % 2 == 0  {
        isBlack = false
      } else if i % 2 != 0 && j % 2 != 0 {
        isBlack = false
      }

      squares = append(squares, Square{Num: i * 8 + j, IsBlack: isBlack, IsActive: false})
    }
  }

  return &Data{Squares: squares, ActiveSquares: []*Square{}}
}

func SetActive(data *Data, num int) {
  square := &(*data).Squares[num]
  (*square).IsActive = true
  data.ActiveSquares = append(data.ActiveSquares, square)
}

func ClearAllActiveSquares(data *Data) {
  for _, square := range (*data).ActiveSquares {
    (*square).IsActive = false
  }
  data.ActiveSquares = []*Square{}
}
