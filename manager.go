package main

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type Manager struct {
	employees []Employee
}

func (m *Manager) AddEmployee(e Employee) {
	m.employees = append(m.employees, e)
}

func (m *Manager) RemoveEmployee(id int) {
	for i, e := range m.employees {
		if e.ID == id {
			m.employees = append(m.employees[:i], m.employees[i+1:]...)
			return
		}
	}
}

func (m *Manager) GetAverageSalary() float64 {
	if len(m.employees) == 0 {
		return 0
	}
	var total float64
	for _, e := range m.employees {
		total += e.Salary
	}
	return total / float64(len(m.employees))
}

func (m *Manager) FindEmployeeByID(id int) *Employee {
	for i := range m.employees {
		if m.employees[i].ID == id {
			return &m.employees[i]
		}
	}
	return nil
}
