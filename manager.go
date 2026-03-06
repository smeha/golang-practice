package main

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type Manager struct {
	Employees []Employee
}

func (m *Manager) AddEmployee(e Employee) {
	m.Employees = append(m.Employees, e)
}

func (m *Manager) RemoveEmployee(id int) {
	for i, e := range m.Employees {
		if e.ID == id {
			m.Employees = append(m.Employees[:i], m.Employees[i+1:]...)
			return
		}
	}
}

func (m *Manager) GetAverageSalary() float64 {
	if len(m.Employees) == 0 {
		return 0
	}
	var total float64
	for _, e := range m.Employees {
		total += e.Salary
	}
	return total / float64(len(m.Employees))
}

func (m *Manager) FindEmployeeByID(id int) *Employee {
	for i := range m.Employees {
		if m.Employees[i].ID == id {
			return &m.Employees[i]
		}
	}
	return nil
}
