package widget

import (
	"fmt"
	"strings"

	"html/template"
)

type DataTableWidget struct {
	Base
	model *TableModel
	style *DataTableStyle
	tmpl *template.Template
}

type HeaderRow struct {
	Labels []string
}

type DataRow struct {
	Items []string
}

type TableModel struct {
	Model
	headerRow *HeaderRow
	dataRows  []*DataRow
}

func NewTableModel(header *HeaderRow, rows []*DataRow) *TableModel {
	return &TableModel{
		Model:     NewModel(),
		headerRow: header,
		dataRows:  rows,
	}
}

func (m *TableModel) Header() *HeaderRow {
	return m.headerRow
}

func (m *TableModel) SetHeader(header *HeaderRow) {
	m.headerRow = header
	m.Notify()
}

func (m *TableModel) Rows() []*DataRow {
	return m.dataRows
}

func (m *TableModel) SetRows(rows []*DataRow) {
	m.dataRows = rows
	m.Notify()
}

func (m *TableModel) AddRow(row *DataRow) {
	m.dataRows = append(m.dataRows, row)
	m.Notify()
}

func DataTable(m *TableModel) *DataTableWidget {
	tmpl, _ := template.New("data_table").Parse(`
  <bx-table>
    <bx-table-head>
      <bx-table-header-row>
	    {{ range .Header.Labels }}
        <bx-table-header-cell>{{.}}</bx-table-header-cell>
		{{ end }}
      </bx-table-header-row>
    </bx-table-head>
    <bx-table-body>
	  {{ range .Rows }}
      <bx-table-row>
	    {{ range .Items }}
        <bx-table-cell>{{.}}</bx-table-cell>
		{{ end }}
      </bx-table-row>
	  {{ end }}
    </bx-table-body>
  </bx-table>
	`)
	dt := &DataTableWidget{
		Base:  NewBase(),
		model: m,
		style: &DataTableStyle{},
		tmpl: tmpl,
	}
	m.AddListener(dt)
	dt.Base.SetWidget(dt)
	return dt
}

func (dt *DataTableWidget) View() string {
	return fmt.Sprintf(
		`<bx-data-table id="%d" %s style="%s">%s</bx-data-table>`,
		dt.ID(),
		dt.style,
		dt.SizeStyle(),
		dt.body(),
	)
}

type TableData struct  {
	Header *HeaderRow
	Rows []*DataRow
}

func (dt *DataTableWidget) body() *strings.Builder {
	var b strings.Builder
	data := &TableData{
		Header: dt.model.headerRow,
		Rows: dt.model.dataRows,
	}
	dt.tmpl.Execute(&b, data)
	return &b
}

func (dt *DataTableWidget) SetStyle(style *DataTableStyle) *DataTableWidget {
	dt.style = style
	return dt
}

type DataTableStyle struct{}

func (s *DataTableStyle) String() string {
	return fmt.Sprintf(``)
}
