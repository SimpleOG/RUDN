package db

import (
	"context"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
	"sync"
)

const (
	//тупо счётчик
	//Образовательная программа
	//educationForm = iota // Форма обучения
	//lvlOp                // Уровень ОП
	theCodeOfTheOopRudn = iota + 2 // Шифр ООП РУДН
	direction_code                 // Код направления
	name_of_the_program            // Наименование программы
	// Дисциплина или вид учебной  работы
	block
	component
	n_v_RUP
	dop_info
	name_of_the_discipline_or_type_of_academic_work
	semester_or_Module
	weeks_per_semester_module
	type_of_educational_work
	lecture_hours
	laboratories_hours
	practise_hours
	type_of_PA_or_GIA
	course_works
	course_projects
	course_Uch_ave_ZE_on_RUP
	pr_ZE_on_RUP
	NIR_ZE_by_RUP
	//контингент обучающихся
	code
	group_number
	of_groups
	subgroups
	total_people
	RF
	foreign
	standard
	calculated
	PK
	//Сведения о ППС
	department
	post
	terms_of_attraction
	full_name
	a_special_feature
	//Объём учебной работы ППС
	lectures
	practice_or_Seminars
	Lab_works_or_Clinical_classes
	current_control
	interim_certification_PO_for_BRS
	registration_of_PA_results
	ongoing_consultations_on_the_discipline
	courseWorks
	courseProjects
	educational_practice
	proc_pedagogical_and_pre_graduate_practices
	NIR
	practices_including_research_of_digital_magistracies
	reviewing_the_abstracts_of_graduate_students
	candidates_exam
	scientific_guidance
	the_leadership_of_the_WRC_or_the_NKR
	review_of_the_WRC
	GEK
	total
	start = 5
	end   = 1645
)

func ReadExcelColumn(find int) ([]string, error) {

	f, err := excelize.OpenFile("считать.xlsx")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	cols, err := f.GetCols("УН сводная")

	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	rows := make([]string, 0)
	for i := start; i < end; i++ {
		rows = append(rows, cols[find][i])
	}

	return rows, nil
}

// подтянем из экселя все штуки( шучу, не все)

func (q *Queries) ReadEducationalProgram() (err error) {
	var theCodeOfTheOOPRUDN []string
	var directionCode []string
	var nameOfTheProgram []string
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		theCodeOfTheOOPRUDN, err = ReadExcelColumn(theCodeOfTheOopRudn)
		if err != nil {
			log.Fatalln("Не удалось найти фио", err)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		directionCode, err = ReadExcelColumn(direction_code)
		if err != nil {
			log.Fatalln("Не удалось найти фио", err)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		nameOfTheProgram, err = ReadExcelColumn(name_of_the_program)
		if err != nil {
			log.Fatalln("Не удалось найти фио", err)
		}
		wg.Done()
	}()
	wg.Wait()
	for i := 0; i < len(nameOfTheProgram); i++ {
		wg.Add(1)
		go func(i int) {
			arg := Create_EducationalProgramParams{
				theCodeOfTheOOPRUDN[i],
				directionCode[i],
				nameOfTheProgram[i],
			}
			if arg.NameOfTheProgram == "" || arg.DirectionCode == "" || arg.TheCodeOfTheOOPRUDN == "" {
				fmt.Println(i)
			}
			_, err = q.Create_EducationalProgram(context.Background(), arg)
			wg.Done()
		}(i)

	}
	wg.Wait()
	return
}
func (q *Queries) ReadDisciplineOrTypeOfAcademicWork() (err error) {
	var Block []string
	var Component []string
	var Nvrup []string
	var Nameof []string
	var Dopinfo []string
	var wg sync.WaitGroup
	wg.Add(5)
	go func() {
		Block, err = ReadExcelColumn(block)
		if err != nil {
			log.Fatalln("Проблема", err)
			return
		}
		wg.Done()
	}()
	go func() {
		Component, err = ReadExcelColumn(component)
		if err != nil {
			log.Fatalln("Проблема", err)
			return
		}
		wg.Done()
	}()
	go func() {
		Nvrup, err = ReadExcelColumn(n_v_RUP)
		if err != nil {
			log.Fatalln("Проблема", err)
			return
		}
		wg.Done()
	}()
	go func() {
		Nameof, err = ReadExcelColumn(name_of_the_discipline_or_type_of_academic_work)
		if err != nil {
			log.Fatalln("Проблема", err)
			return
		}
		wg.Done()
	}()
	go func() {
		Dopinfo, err = ReadExcelColumn(dop_info)
		if err != nil {
			log.Fatalln("Проблема", err)
			return
		}
		wg.Done()
	}()
	wg.Wait()
	for i := 0; i < len(Nameof); i++ {
		wg.Add(1)
		go func(i int) {
			arg := Create_Discipline_or_type_of_academic_workParams{
				Block:                                   Block[i],
				Component:                               Component[i],
				NVRUP:                                   Nvrup[i],
				NameOfTheDisciplineOrTypeOfAcademicWork: Nameof[i],
				DopInfo:                                 Dopinfo[i],
			}
			_, err = q.Create_Discipline_or_type_of_academic_work(context.Background(), arg)
			if err != nil {
				return
			}
			wg.Done()

		}(i)
	}
	wg.Wait()
	return
}
func (q *Queries) ReadKW() (err error) {
	var SemesterOrModule []string
	var WeeksPerSemesterModule []string
	var TypeOfEducationalWork []string
	var LectureHours []string
	var LaboratoriesHours []string
	var PractiseHours []string
	var TypeOfPaOrGia []string
	var CourseWorks []string
	var CourseProjects []string
	var CourseUchAveZeOnRup []string
	var PrZeOnRup []string
	var nirZeByRup []string
	var wg sync.WaitGroup
	wg.Add(12)
	go func() {
		SemesterOrModule, err = ReadExcelColumn(semester_or_Module)
		if err != nil {
			log.Fatalln("Проблема", err)
		}
		wg.Done()
	}()
	go func() {
		WeeksPerSemesterModule, err = ReadExcelColumn(weeks_per_semester_module)
		if err != nil {
			log.Fatalln("Проблема", err)
		}
		wg.Done()
	}()
	go func() {
		TypeOfEducationalWork, err = ReadExcelColumn(type_of_educational_work)
		if err != nil {
			log.Fatalln("Проблема", err)
		}
		wg.Done()
	}()
	go func() {
		LectureHours, err = ReadExcelColumn(lecture_hours)
		if err != nil {
			log.Fatalln("Проблема", err)
		}
		wg.Done()
	}()
	go func() {
		LaboratoriesHours, err = ReadExcelColumn(laboratories_hours)
		if err != nil {
			log.Fatalln("Проблема", err)
		}
		wg.Done()
	}()
	go func() {
		PractiseHours, err = ReadExcelColumn(practise_hours)
		if err != nil {
			log.Fatalln("Проблема", err)
		}
		wg.Done()
	}()
	go func() {
		TypeOfPaOrGia, err = ReadExcelColumn(type_of_PA_or_GIA)
		if err != nil {
			log.Fatalln("Проблема", err)
		}
		wg.Done()
	}()
	go func() {
		CourseWorks, err = ReadExcelColumn(course_works)
		if err != nil {
			log.Fatalln("Проблема", err)
		}
		wg.Done()
	}()
	go func() {
		CourseProjects, err = ReadExcelColumn(course_projects)
		if err != nil {
			log.Fatalln("Проблема", err)
		}
		wg.Done()
	}()
	go func() {
		CourseUchAveZeOnRup, err = ReadExcelColumn(course_Uch_ave_ZE_on_RUP)
		if err != nil {
			log.Fatalln("Проблема", err)
		}
		wg.Done()
	}()
	go func() {
		PrZeOnRup, err = ReadExcelColumn(pr_ZE_on_RUP)
		if err != nil {
			log.Fatalln("Проблема", err)
		}
		wg.Done()
	}()
	go func() {
		nirZeByRup, err = ReadExcelColumn(NIR_ZE_by_RUP)
		if err != nil {
			log.Fatalln("Проблема", err)
		}
		wg.Done()
	}()
	wg.Wait()
	var (
		weeks        int
		lecture      int
		laboratories int
		practise     int
	)
	for i := 0; i < len(SemesterOrModule); i++ {
		wg.Add(1)
		go func(i int) {
			if WeeksPerSemesterModule[i] == "" {
				weeks = 0
			} else {
				weeks, err = strconv.Atoi(WeeksPerSemesterModule[i])
				if err != nil {
					log.Fatalln(err)
				}
			}
			if LectureHours[i] == "" {
				lecture = 0
			} else {
				lecture, err = strconv.Atoi(LectureHours[i])
				if err != nil {
					log.Fatalln(err)
				}
			}
			if LaboratoriesHours[i] == "" {
				laboratories = 0
			} else {
				laboratories, err = strconv.Atoi(LaboratoriesHours[i])
				if err != nil {
					log.Fatalln(err)
				}
			}
			if PractiseHours[i] == "" {
				practise = 0
			} else {
				practise, err = strconv.Atoi(PractiseHours[i])
				if err != nil {
					log.Fatalln(err)
				}
			}
			arg := Create_k_wParams{
				SemesterOrModule:       SemesterOrModule[i],
				WeeksPerSemesterModule: int32(weeks),
				TypeOfEducationalWork:  TypeOfEducationalWork[i],
				LectureHours:           int32(lecture),
				LaboratoriesHours:      int32(laboratories),
				PractiseHours:          int32(practise),
				TypeOfPAOrGIA:          TypeOfPaOrGia[i],
				CourseWorks:            CourseWorks[i],
				CourseProjects:         CourseProjects[i],
				CourseUchAveZEOnRUP:    CourseUchAveZeOnRup[i],
				PrZEOnRUP:              PrZeOnRup[i],
				NIRZEByRUP:             nirZeByRup[i],
			}
			_, err = q.Create_k_w(context.Background(), arg)
			if err != nil {
				log.Fatalln(err)
			}
			wg.Done()
		}(i)

	}
	wg.Wait()
	return
}
func (q *Queries) ReadTheContingentOfStudents() (err error) {
	var Code []string
	var groupNumber []string
	var ofGroups []string
	var Subgroups []string
	var totalPeople []string
	var rf []string
	var Foreign []string
	var standart []string
	var Calculated []string
	var pk []string

	var wg sync.WaitGroup
	wg.Add(10)
	go func() {
		Code, err = ReadExcelColumn(code)
		if err != nil {
			log.Fatalln(err)
		}
		wg.Done()
	}()
	go func() {
		groupNumber, err = ReadExcelColumn(group_number)
		if err != nil {
			log.Fatalln(err)
		}
		wg.Done()
	}()
	go func() {
		ofGroups, err = ReadExcelColumn(of_groups)
		if err != nil {
			log.Fatalln(err)
		}
		wg.Done()
	}()
	go func() {
		Subgroups, err = ReadExcelColumn(subgroups)
		if err != nil {
			log.Fatalln(err)
		}
		wg.Done()
	}()
	go func() {
		totalPeople, err = ReadExcelColumn(total_people)
		if err != nil {
			log.Fatalln(err)
		}
		wg.Done()
	}()
	go func() {
		rf, err = ReadExcelColumn(RF)
		if err != nil {
			log.Fatalln(err)
		}
		wg.Done()
	}()
	go func() {
		Foreign, err = ReadExcelColumn(foreign)
		if err != nil {
			log.Fatalln(err)
		}
		wg.Done()
	}()
	go func() {
		standart, err = ReadExcelColumn(standard)
		if err != nil {
			log.Fatalln(err)
		}
		wg.Done()
	}()
	go func() {
		Calculated, err = ReadExcelColumn(calculated)
		if err != nil {
			log.Fatalln(err)
		}
		wg.Done()
	}()
	go func() {
		pk, err = ReadExcelColumn(PK)
		if err != nil {
			log.Fatalln(err)
		}
		wg.Done()
	}()

	wg.Wait()

	for i := 0; i < len(groupNumber); i++ {
		wg.Add(1)
		go func(i int) {
			arg := Create_the_contingent_of_studentsParams{
				Code:        Code[i],
				GroupNumber: groupNumber[i],
				OfGroups:    ofGroups[i],
				Subgroups:   Subgroups[i],
				TotalPeople: totalPeople[i],
				RF:          rf[i],
				Foreign:     Foreign[i],
				Standard:    standart[i],
				Calculated:  Calculated[i],
				PK:          pk[i],
			}
			_, err = q.Create_the_contingent_of_students(context.Background(), arg)
			if err != nil {
				log.Fatalln(err)
			}
			wg.Done()
		}(i)

	}
	wg.Wait()
	return
}
func (q *Queries) ReadInformationAboutPps() (err error) {
	var Department []string
	var Post []string
	var TermsOfAttraction []string
	var FullName []string
	var ASpecialFeature []string
	var wg sync.WaitGroup
	wg.Add(5)
	go func() {
		Department, err = ReadExcelColumn(department)
		wg.Done()
	}()
	go func() {
		Post, err = ReadExcelColumn(post)
		wg.Done()
	}()
	go func() {
		TermsOfAttraction, err = ReadExcelColumn(terms_of_attraction)
		wg.Done()
	}()
	go func() {
		FullName, err = ReadExcelColumn(full_name)
		wg.Done()
	}()
	go func() {
		ASpecialFeature, err = ReadExcelColumn(a_special_feature)
		wg.Done()
	}()
	wg.Wait()
	for i := 0; i < len(FullName); i++ {
		wg.Add(1)
		go func(i int) {
			arg := Create_information_about_PPSParams{
				Department:        Department[i],
				Post:              Post[i],
				TermsOfAttraction: TermsOfAttraction[i],
				FullName:          FullName[i],
				ASpecialFeature:   ASpecialFeature[i],
			}
			_, err = q.Create_information_about_PPS(context.Background(), arg)
			wg.Done()
		}(i)
	}
	wg.Wait()
	return
}
func (q *Queries) ReadTheAmountOfTeachingWorkOfTheTeachingStaff() (err error) {
	var Lectures []string
	var PracticeOrSeminars []string
	var LabWorksOrClinicalClasses []string
	var CurrentControl []string
	var InterimCertificationPoForBrs []string
	var RegistrationOfPaResults []string
	var OngoingConsultationsOnTheDiscipline []string
	var CourseWorks []string
	var CourseProjects []string
	var EducationalPractice []string
	var ProcPedagogicalAndPreGraduatePractices []string
	var nir []string
	var PracticesIncludingResearchOfDigitalMagistracies []string
	var ReviewingTheAbstractsOfGraduateStudents []string
	var CandidatesExam []string
	var ScientificGuidance []string
	var TheLeadershipOfTheWrcOrTheNkr []string
	var ReviewOfTheWrc []string
	var gek []string
	var Total []string
	var wg sync.WaitGroup
	wg.Add(20)
	go func() {
		Lectures, err = ReadExcelColumn(lectures)
		wg.Done()
	}()
	go func() {
		PracticeOrSeminars, err = ReadExcelColumn(practice_or_Seminars)
		wg.Done()
	}()
	go func() {
		LabWorksOrClinicalClasses, err = ReadExcelColumn(Lab_works_or_Clinical_classes)
		wg.Done()
	}()
	go func() {
		CurrentControl, err = ReadExcelColumn(current_control)
		wg.Done()
	}()
	go func() {
		InterimCertificationPoForBrs, err = ReadExcelColumn(interim_certification_PO_for_BRS)
		wg.Done()
	}()
	go func() {
		RegistrationOfPaResults, err = ReadExcelColumn(registration_of_PA_results)
		wg.Done()
	}()
	go func() {
		OngoingConsultationsOnTheDiscipline, err = ReadExcelColumn(ongoing_consultations_on_the_discipline)
		wg.Done()
	}()
	go func() {
		CourseWorks, err = ReadExcelColumn(courseWorks)
		wg.Done()
	}()
	go func() {
		CourseProjects, err = ReadExcelColumn(courseProjects)
		wg.Done()
	}()
	go func() {
		EducationalPractice, err = ReadExcelColumn(educational_practice)
		wg.Done()
	}()
	go func() {
		ProcPedagogicalAndPreGraduatePractices, err = ReadExcelColumn(proc_pedagogical_and_pre_graduate_practices)
		wg.Done()
	}()
	go func() {
		nir, err = ReadExcelColumn(NIR)
		wg.Done()
	}()
	go func() {
		PracticesIncludingResearchOfDigitalMagistracies, err = ReadExcelColumn(practices_including_research_of_digital_magistracies)
		wg.Done()
	}()
	go func() {
		ReviewingTheAbstractsOfGraduateStudents, err = ReadExcelColumn(reviewing_the_abstracts_of_graduate_students)
		wg.Done()
	}()
	go func() {
		CandidatesExam, err = ReadExcelColumn(candidates_exam)
		wg.Done()
	}()
	go func() {
		ScientificGuidance, err = ReadExcelColumn(scientific_guidance)
		wg.Done()
	}()
	go func() {
		TheLeadershipOfTheWrcOrTheNkr, err = ReadExcelColumn(the_leadership_of_the_WRC_or_the_NKR)
		wg.Done()
	}()
	go func() {
		ReviewOfTheWrc, err = ReadExcelColumn(review_of_the_WRC)
		wg.Done()
	}()
	go func() {
		gek, err = ReadExcelColumn(GEK)
		wg.Done()
	}()
	go func() {
		Total, err = ReadExcelColumn(total)
		wg.Done()
	}()
	wg.Wait()
	for i := 0; i < len(Lectures); i++ {
		wg.Add(1)
		go func(i int) {
			arg := Create_the_amount_of_teaching_work_of_the_teaching_staffParams{
				Lectures:                               Lectures[i],
				PracticeOrSeminars:                     PracticeOrSeminars[i],
				LabWorksOrClinicalClasses:              LabWorksOrClinicalClasses[i],
				CurrentControl:                         CurrentControl[i],
				InterimCertificationPOForBRS:           InterimCertificationPoForBrs[i],
				RegistrationOfPAResults:                RegistrationOfPaResults[i],
				OngoingConsultationsOnTheDiscipline:    OngoingConsultationsOnTheDiscipline[i],
				CourseWorks:                            CourseWorks[i],
				CourseProjects:                         CourseProjects[i],
				EducationalPractice:                    EducationalPractice[i],
				ProcPedagogicalAndPreGraduatePractices: ProcPedagogicalAndPreGraduatePractices[i],
				NIR:                                    nir[i],
				PracticesIncludingResearchOfDigitalMagistracies: PracticesIncludingResearchOfDigitalMagistracies[i],
				ReviewingTheAbstractsOfGraduateStudents:         ReviewingTheAbstractsOfGraduateStudents[i],
				CandidatesExam:                                  CandidatesExam[i],
				ScientificGuidance:                              ScientificGuidance[i],
				TheLeadershipOfTheWRCOrTheNKR:                   TheLeadershipOfTheWrcOrTheNkr[i],
				ReviewOfTheWRC:                                  ReviewOfTheWrc[i],
				GEK:                                             gek[i],
				Total:                                           Total[i],
			}
			_, err = q.Create_the_amount_of_teaching_work_of_the_teaching_staff(context.Background(), arg)
			if err != nil {
				log.Fatalln(err)
			}
			wg.Done()
		}(i)

	}
	wg.Wait()
	return
}

func (q *Queries) MakeItAll() (err error) {
	wg := new(sync.WaitGroup)
	wg.Add(6)
	go func() {
		err = q.ReadEducationalProgram()
		if err != nil {
			log.Fatalln(err)
		}
		wg.Done()
	}()
	go func() {
		err = q.ReadDisciplineOrTypeOfAcademicWork()
		if err != nil {
			log.Fatalln(err)
		}
		wg.Done()
		fmt.Println("Я всё")
	}()
	go func() {
		err = q.ReadKW()
		if err != nil {
			log.Fatalln(err)
		}
		wg.Done()
	}()
	go func() {
		err = q.ReadTheContingentOfStudents()
		if err != nil {
			log.Fatalln(err)
		}
		wg.Done()
	}()
	go func() {
		err = q.ReadInformationAboutPps()
		if err != nil {
			log.Fatalln(err)
		}
		wg.Done()
	}()
	go func() {
		err = q.ReadTheAmountOfTeachingWorkOfTheTeachingStaff()
		if err != nil {
			log.Fatalln(err)
		}
		wg.Done()
	}()
	wg.Wait()
	return
}
