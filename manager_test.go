package main

import "testing"

func TestManagerAddAndFind(t *testing.T) {
	m := Manager{}
	m.AddEmployee(Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})

	found := m.FindEmployeeByID(1)
	if found == nil {
		t.Fatal("expected to find employee with ID 1")
	}
	if found.Name != "Alice" {
		t.Errorf("got name %q, want %q", found.Name, "Alice")
	}
}

func TestManagerRemoveEmployee(t *testing.T) {
	m := Manager{}
	m.AddEmployee(Employee{ID: 1, Name: "Alice", Salary: 70000})
	m.AddEmployee(Employee{ID: 2, Name: "Bob", Salary: 60000})
	m.RemoveEmployee(1)

	if m.FindEmployeeByID(1) != nil {
		t.Error("expected employee 1 to be removed")
	}
	if m.FindEmployeeByID(2) == nil {
		t.Error("expected employee 2 to still exist")
	}
}

func TestManagerRemoveNonExistent(t *testing.T) {
	m := Manager{}
	m.AddEmployee(Employee{ID: 1, Name: "Alice", Salary: 70000})
	m.RemoveEmployee(99) // should not panic or remove anything
	if m.FindEmployeeByID(1) == nil {
		t.Error("expected employee 1 to still exist after removing non-existent ID")
	}
}

func TestManagerGetAverageSalary(t *testing.T) {
	tests := []struct {
		employees []Employee
		want      float64
	}{
		{[]Employee{}, 0},
		{[]Employee{{ID: 1, Salary: 100}}, 100},
		{[]Employee{{ID: 1, Salary: 100}, {ID: 2, Salary: 200}}, 150},
	}
	for _, tt := range tests {
		m := Manager{}
		for _, e := range tt.employees {
			m.AddEmployee(e)
		}
		if got := m.GetAverageSalary(); got != tt.want {
			t.Errorf("GetAverageSalary() = %v, want %v", got, tt.want)
		}
	}
}

func TestManagerFindReturnsPointerIntoSlice(t *testing.T) {
	m := Manager{}
	m.AddEmployee(Employee{ID: 1, Name: "Alice", Salary: 70000})

	e := m.FindEmployeeByID(1)
	e.Name = "Updated"

	if m.FindEmployeeByID(1).Name != "Updated" {
		t.Error("FindEmployeeByID should return a pointer into the slice, not a copy")
	}
}
