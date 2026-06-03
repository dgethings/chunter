package document

func (d *Document) OffsetAt(pos Position) int {
	line := int(pos.Line)
	char := int(pos.Character)
	if line < 0 || line >= len(d.Lines) {
		return -1
	}
	offset := 0
	for i := 0; i < line; i++ {
		offset += len(d.Lines[i]) + 1
	}
	offset += char
	return offset
}

type Position struct {
	Line      uint
	Character uint
}
