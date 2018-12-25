package number

func ToInt(i iterface{}) int64 {
  return reflect.ValueOf(i).Int()
}
