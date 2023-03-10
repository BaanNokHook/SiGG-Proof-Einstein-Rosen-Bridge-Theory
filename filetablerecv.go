package main

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import "github.com/therecipe/qt/core"

const (
	Filename = int(core.Qt__UserRole) + 1<<iota
	Filesize
	Transmitted
	Status
)

type TableItem struct {
	filename    string
	filesize    string
	transmitted string
	status      string
}

type RecvFileTableModel struct {
	core.QAbstractTableModel

	_ func() `constructor:"init"`

	_ func()                                                                                  `signal:"remove,auto"`
	_ func(item []*core.QVariant)                                                             `signal:"add,auto"`
	_ func(filename string, filesize string, transmitted string, status string)               `signal:"editLast,auto"`
	_ func(tableIdx int, filename string, filesize string, transmitted string, status string) `signal:"editIdx,auto"`

	modelData []TableItem
}

func (m *RecvFileTableModel) init() {
	m.modelData = []TableItem{
		//		{"test1.txt", "1000", "1000", "Done"},
		//		{"test2.txt", "1000", "0", "Started"},
	}

	m.ConnectRoleNames(m.roleNames)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectData(m.data)
}

func (m *RecvFileTableModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		Filename:    core.NewQByteArray2("Filename", -1),
		Filesize:    core.NewQByteArray2("Filesize", -1),
		Transmitted: core.NewQByteArray2("Transmitted", -1),
		Status:      core.NewQByteArray2("Status", -1),
	}
}

func (m *RecvFileTableModel) rowCount(*core.QModelIndex) int {
	return len(m.modelData)
}

func (m *RecvFileTableModel) columnCount(*core.QModelIndex) int {
	return 4
}

func (m *RecvFileTableModel) data(index *core.QModelIndex, role int) *core.QVariant {
	item := m.modelData[index.Row()]
	switch role {
	case Filename:
		return core.NewQVariant1(item.filename)
	case Filesize:
		return core.NewQVariant1(item.filesize)
	case Transmitted:
		return core.NewQVariant1(item.transmitted)
	case Status:
		return core.NewQVariant1(item.status)

	}
	return core.NewQVariant()
}

func (m *RecvFileTableModel) remove() {
	if len(m.modelData) == 0 {
		return
	}
	m.BeginRemoveRows(core.NewQModelIndex(), len(m.modelData)-1, len(m.modelData)-1)
	m.modelData = m.modelData[:len(m.modelData)-1]
	m.EndRemoveRows()
}

func (m *RecvFileTableModel) add(item []*core.QVariant) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	m.modelData = append(
		m.modelData,
		TableItem{
			item[0].ToString(),
			item[1].ToString(),
			item[2].ToString(),
			item[3].ToString(),
		})
	m.EndInsertRows()
}

func (m *RecvFileTableModel) addNative(a string, b string, c string, d string) int {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	m.modelData = append(
		m.modelData,
		TableItem{
			a,
			b,
			c,
			d,
		})
	m.EndInsertRows()

	return len(m.modelData) - 1
}

func (m *RecvFileTableModel) editLast(filename string, filesize string, transmitted string, status string) {
	if len(m.modelData) == 0 {
		return
	}
	m.modelData[len(m.modelData)-1] = TableItem{filename, filesize, transmitted, status}
	m.DataChanged(
		m.Index(len(m.modelData)-1, 0, core.NewQModelIndex()),
		m.Index(len(m.modelData)-1, 1, core.NewQModelIndex()),
		[]int{Filename, Filesize, Transmitted, Status})
}

func (m *RecvFileTableModel) editIdx(tableIdx int, filename string, filesize string, transmitted string, status string) {
	if len(m.modelData) == 0 {
		return
	}
	m.modelData[tableIdx] = TableItem{filename, filesize, transmitted, status}
	m.DataChanged(
		m.Index(tableIdx, 0, core.NewQModelIndex()),
		m.Index(tableIdx, 1, core.NewQModelIndex()),
		[]int{Filename, Filesize, Transmitted, Status})
}
