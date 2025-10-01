package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	assignmentGrade := 100
	examGrade := 100
	essayGrade := 100

	t.Logf("Testing with %d%% %s, %d%% %s, %d%% %s", assignmentGrade, Assignment, examGrade, Exam, essayGrade, Essay)

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", assignmentGrade, Assignment)
	gradeCalculator.AddGrade("exam 1", examGrade, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", essayGrade, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	assignmentGrade := 80
	examGrade := 81
	essayGrade := 85

	t.Logf("Testing with %d%% %s, %d%% %s, %d%% %s", assignmentGrade, Assignment, examGrade, Exam, essayGrade, Essay)

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", assignmentGrade, Assignment)
	gradeCalculator.AddGrade("exam 1", examGrade, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", essayGrade, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	assignmentGrade := 70
	examGrade := 71
	essayGrade := 75

	t.Logf("Testing with %d%% %s, %d%% %s, %d%% %s", assignmentGrade, Assignment, examGrade, Exam, essayGrade, Essay)

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", assignmentGrade, Assignment)
	gradeCalculator.AddGrade("exam 1", examGrade, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", essayGrade, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	assignmentGrade := 60
	examGrade := 61
	essayGrade := 65

	t.Logf("Testing with %d%% %s, %d%% %s, %d%% %s", assignmentGrade, Assignment, examGrade, Exam, essayGrade, Essay)

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", assignmentGrade, Assignment)
	gradeCalculator.AddGrade("exam 1", examGrade, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", essayGrade, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	assignmentGrade := 50
	examGrade := 51
	essayGrade := 55

	t.Logf("Testing with %d%% %s, %d%% %s, %d%% %s", assignmentGrade, Assignment, examGrade, Exam, essayGrade, Essay)

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", assignmentGrade, Assignment)
	gradeCalculator.AddGrade("exam 1", examGrade, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", essayGrade, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
