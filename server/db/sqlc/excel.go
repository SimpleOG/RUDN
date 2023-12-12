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
	size  = end - start
)

func ReadExcel() ([][]string, error) {
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
		log.Fatalln(err)
	}
	for i := range cols {
		cols[i] = cols[i][start:end]

	}
	return cols, nil
}

// подтянем из экселя все штуки( шучу, не все)

func (q *Queries) ReadEducationalProgram(data [][]string) ([size]EducationalProgram, error) {
	lock := new(sync.Mutex)
	m := make(map[int]EducationalProgram)
	theCodeOfTheOOPRUDN := data[theCodeOfTheOopRudn]
	directionCode := data[direction_code]
	nameOfTheProgram := data[name_of_the_program]
	var wg sync.WaitGroup
	for i := 0; i < size; i++ {
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
			n, err := q.Create_EducationalProgram(context.Background(), arg)
			if err != nil {
				log.Fatalln(err)
			}
			lock.Lock()
			m[i] = n
			lock.Unlock()
			wg.Done()
		}(i)

	}
	wg.Wait()
	var arr [1640]EducationalProgram
	for i, v := range m {
		arr[i] = v
	}
	return arr, nil
}

func (q *Queries) ReadDisciplineOrTypeOfAcademicWork(data [][]string) ([size]DisciplineOrTypeOfAcademicWork, error) {
	var err error
	lock := new(sync.Mutex)
	m := make(map[int]DisciplineOrTypeOfAcademicWork)
	Block := data[block]
	Component := data[component]
	Nvrup := data[n_v_RUP]
	Nameof := data[name_of_the_discipline_or_type_of_academic_work]
	Dopinfo := data[dop_info]
	var wg sync.WaitGroup
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {
			arg := Create_Discipline_or_type_of_academic_workParams{
				Block:                                   Block[i],
				Component:                               Component[i],
				NVRUP:                                   Nvrup[i],
				NameOfTheDisciplineOrTypeOfAcademicWork: Nameof[i],
				DopInfo:                                 Dopinfo[i],
			}
			c, err := q.Create_Discipline_or_type_of_academic_work(context.Background(), arg)
			if err != nil {
				log.Fatalln(err)
			}
			lock.Lock()
			m[i] = c
			lock.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	var arr [1640]DisciplineOrTypeOfAcademicWork
	for i, v := range m {
		arr[i] = v
	}
	return arr, err
}

func (q *Queries) ReadKW(data [][]string) ([size]KW, error) {
	var err error
	lock := new(sync.Mutex)
	m := make(map[int]KW)
	SemesterOrModule := data[semester_or_Module]
	WeeksPerSemesterModule := data[weeks_per_semester_module]
	TypeOfEducationalWork := data[type_of_educational_work]
	LectureHours := data[lecture_hours]
	LaboratoriesHours := data[laboratories_hours]
	PractiseHours := data[practise_hours]
	TypeOfPaOrGia := data[type_of_PA_or_GIA]
	CourseWorks := data[courseWorks]
	CourseProjects := data[course_projects]
	CourseUchAveZeOnRup := data[course_Uch_ave_ZE_on_RUP]
	PrZeOnRup := data[pr_ZE_on_RUP]
	nirZeByRup := data[NIR_ZE_by_RUP]
	var wg sync.WaitGroup
	var (
		weeks        int
		lecture      int
		laboratories int
		practise     int
	)
	for i := 0; i < size; i++ {
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
			c, err := q.Create_k_w(context.Background(), arg)
			lock.Lock()
			m[i] = c
			lock.Unlock()
			if err != nil {
				log.Fatalln(err)
			}
			wg.Done()
		}(i)

	}
	wg.Wait()
	var arr [size]KW
	for i, v := range m {
		arr[i] = v
	}
	return arr, err
}

func (q *Queries) ReadTheContingentOfStudents(data [][]string) ([size]TheContingentOfStudent, error) {
	var err error
	lock := new(sync.Mutex)
	m := make(map[int]TheContingentOfStudent)
	Code := data[code]
	groupNumber := data[group_number]
	ofGroups := data[of_groups]
	Subgroups := data[subgroups]
	totalPeople := data[total_people]
	rf := data[RF]
	Foreign := data[foreign]
	standart := data[standard]
	Calculated := data[calculated]
	pk := data[PK]
	var wg sync.WaitGroup

	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {
			arg := Create_the_contingent_of_studentsParams{
				Code:        Code[i],
				GroupName:   Code[i] + groupNumber[i],
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
			c, err := q.Create_the_contingent_of_students(context.Background(), arg)
			if err != nil {
				log.Fatalln(err)
			}
			lock.Lock()
			m[i] = c
			lock.Unlock()
			wg.Done()
		}(i)

	}
	wg.Wait()
	var arr [size]TheContingentOfStudent
	for i, v := range m {
		arr[i] = v
	}
	return arr, err
}

func (q *Queries) ReadInformationAboutPps(data [][]string) ([size]InformationAboutPP, error) {
	var err error
	lock := new(sync.Mutex)
	m := make(map[int]InformationAboutPP)
	Department := data[department]
	Post := data[post]
	TermsOfAttraction := data[terms_of_attraction]
	FullName := data[full_name]
	ASpecialFeature := data[a_special_feature]
	var wg sync.WaitGroup
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
			c, err := q.Create_information_about_PPS(context.Background(), arg)
			if err != nil {
				log.Fatalln(err)
			}
			lock.Lock()
			m[i] = c
			lock.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	var arr [size]InformationAboutPP
	for i, v := range m {
		arr[i] = v
	}
	return arr, err
}

func (q *Queries) ReadTheAmountOfTeachingWorkOfTheTeachingStaff(data [][]string) ([size]TheAmountOfTeachingWorkOfTheTeachingStaff, error) {
	var err error
	lock := new(sync.Mutex)
	m := make(map[int]TheAmountOfTeachingWorkOfTheTeachingStaff)
	Lectures := data[lectures]
	PracticeOrSeminars := data[practice_or_Seminars]
	LabWorksOrClinicalClasses := data[Lab_works_or_Clinical_classes]
	CurrentControl := data[current_control]
	InterimCertificationPoForBrs := data[interim_certification_PO_for_BRS]
	RegistrationOfPaResults := data[registration_of_PA_results]
	OngoingConsultationsOnTheDiscipline := data[ongoing_consultations_on_the_discipline]
	CourseWorks := data[courseWorks]
	CourseProjects := data[courseProjects]
	EducationalPractice := data[educational_practice]
	ProcPedagogicalAndPreGraduatePractices := data[proc_pedagogical_and_pre_graduate_practices]
	nir := data[NIR]
	PracticesIncludingResearchOfDigitalMagistracies := data[practices_including_research_of_digital_magistracies]
	ReviewingTheAbstractsOfGraduateStudents := data[reviewing_the_abstracts_of_graduate_students]
	CandidatesExam := data[candidates_exam]
	ScientificGuidance := data[scientific_guidance]
	TheLeadershipOfTheWrcOrTheNkr := data[the_leadership_of_the_WRC_or_the_NKR]
	ReviewOfTheWrc := data[review_of_the_WRC]
	gek := data[GEK]
	Total := data[total]
	var wg sync.WaitGroup
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {
			to, err := strconv.ParseFloat(Total[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			l, err := strconv.ParseFloat(Lectures[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			p, err := strconv.ParseFloat(PracticeOrSeminars[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			lab, err := strconv.ParseFloat(LabWorksOrClinicalClasses[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			c, err := strconv.ParseFloat(CurrentControl[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			inter, err := strconv.ParseFloat(InterimCertificationPoForBrs[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			reg, err := strconv.ParseFloat(RegistrationOfPaResults[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			ong, err := strconv.ParseFloat(OngoingConsultationsOnTheDiscipline[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			course, err := strconv.ParseFloat(CourseWorks[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			projects, err := strconv.ParseFloat(CourseProjects[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			edu, err := strconv.ParseFloat(EducationalPractice[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			proc, err := strconv.ParseFloat(ProcPedagogicalAndPreGraduatePractices[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			n, err := strconv.ParseFloat(nir[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			prac, err := strconv.ParseFloat(PracticesIncludingResearchOfDigitalMagistracies[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			rev, err := strconv.ParseFloat(ReviewingTheAbstractsOfGraduateStudents[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			cand, err := strconv.ParseFloat(CandidatesExam[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			sci, err := strconv.ParseFloat(ScientificGuidance[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			leader, err := strconv.ParseFloat(TheLeadershipOfTheWrcOrTheNkr[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			wrc, err := strconv.ParseFloat(ReviewOfTheWrc[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			ge, err := strconv.ParseFloat(gek[i], 10)
			if err != nil {
				log.Fatalln(err)
			}
			arg := Create_the_amount_of_teaching_work_of_the_teaching_staffParams{
				Lectures:                               l,
				PracticeOrSeminars:                     p,
				LabWorksOrClinicalClasses:              lab,
				CurrentControl:                         c,
				InterimCertificationPOForBRS:           inter,
				RegistrationOfPAResults:                reg,
				OngoingConsultationsOnTheDiscipline:    ong,
				CourseWorks:                            course,
				CourseProjects:                         projects,
				EducationalPractice:                    edu,
				ProcPedagogicalAndPreGraduatePractices: proc,
				NIR:                                    n,
				PracticesIncludingResearchOfDigitalMagistracies: prac,
				ReviewingTheAbstractsOfGraduateStudents:         rev,
				CandidatesExam:                                  cand,
				ScientificGuidance:                              sci,
				TheLeadershipOfTheWRCOrTheNKR:                   leader,
				ReviewOfTheWRC:                                  wrc,
				GEK:                                             ge,
				Total:                                           to,
			}

			amount, err := q.Create_the_amount_of_teaching_work_of_the_teaching_staff(context.Background(), arg)

			if err != nil {
				log.Fatalln(err)
			}
			lock.Lock()
			m[i] = amount
			lock.Unlock()
			wg.Done()
		}(i)

	}
	wg.Wait()
	var arr [size]TheAmountOfTeachingWorkOfTheTeachingStaff
	for i, v := range m {
		arr[i] = v
	}
	return arr, err
}

func (q *Queries) ReadItAll() error {
	var err error
	data, err := ReadExcel()
	if err != nil {
		return err
	}
	log.Println("Всего прочитано элементов ", len(data)*end)
	wg := new(sync.WaitGroup)
	var program [size]EducationalProgram
	var discipline [size]DisciplineOrTypeOfAcademicWork
	var kw [size]KW
	var group [size]TheContingentOfStudent
	var pps [size]InformationAboutPP
	var amount [size]TheAmountOfTeachingWorkOfTheTeachingStaff
	wg.Add(1)
	go func() {
		wg.Add(6)
		go func() {
			program, err = q.ReadEducationalProgram(data)
			if err != nil {
				log.Fatalln(err)
			}
			wg.Done()
		}()
		go func() {
			discipline, err = q.ReadDisciplineOrTypeOfAcademicWork(data)
			if err != nil {
				log.Fatalln(err)
			}
			wg.Done()
		}()
		go func() {
			kw, err = q.ReadKW(data)
			if err != nil {
				log.Fatalln(err)
			}
			wg.Done()
		}()
		go func() {
			group, err = q.ReadTheContingentOfStudents(data)
			if err != nil {
				log.Fatalln(err)
			}
			wg.Done()
		}()
		go func() {
			pps, err = q.ReadInformationAboutPps(data)
			if err != nil {
				log.Fatalln(err)
			}
			wg.Done()
		}()
		go func() {
			amount, err = q.ReadTheAmountOfTeachingWorkOfTheTeachingStaff(data)
			if err != nil {
				log.Fatalln(err)
			}
			wg.Done()
		}()
		wg.Done()
	}()
	wg.Wait()
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {
			arg := Create_togetherParams{
				ProgramName:    program[i].NameOfTheProgram,
				DisciplineName: discipline[i].NameOfTheDisciplineOrTypeOfAcademicWork,
				GroupName:      group[i].GroupName,
				TeacherName:    pps[i].FullName,
				KWID:           kw[i].ID,
				AmountID:       amount[i].ID,
			}
			wg.Add(2)
			go func() {
				//go func() {
				//	_, err := q.Create_Program_group(context.Background(), Create_Program_groupParams{
				//		ProgramName: arg.ProgramName,
				//		GroupName:   arg.GroupName,
				//	})
				//	if err != nil {
				//		log.Fatalln("строка 449 ", err)
				//	}
				//	wg.Done()
				//}()
				//go func() {
				//	_, err = q.Create_discipline_group(context.Background(), Create_discipline_groupParams{
				//		DisciplineName: arg.DisciplineName,
				//		GroupName:      arg.GroupName,
				//	})
				//	if err != nil {
				//		log.Fatalln("строка 509 ", err)
				//	}
				//	wg.Done()
				//}()
				//go func() {
				//	_, err = q.Create_teacher_group(context.Background(), Create_teacher_groupParams{
				//		TeacherName:    arg.TeacherName,
				//		DisciplineName: arg.DisciplineName,
				//		GroupName:      arg.GroupName,
				//	})
				//	if err != nil {
				//		log.Fatalln("строка 519 ", err)
				//	}
				//	wg.Done()
				//}()
				//go func() {
				//	_, err = q.Create_group_kw(context.Background(), Create_group_kwParams{
				//		KwID:           arg.KWID,
				//		GroupName:      arg.GroupName,
				//		DisciplineName: arg.DisciplineName,
				//	})
				//	if err != nil {
				//		log.Fatalln("строка 530 ", err)
				//	}
				//	wg.Done()
				//}()
				//go func() {
				//	_, err = q.Create_group_hours_discipline(context.Background(), Create_group_hours_disciplineParams{
				//		GroupName:      arg.GroupName,
				//		DisciplineName: arg.DisciplineName,
				//		AmountID:       arg.AmountID,
				//	})
				//	if err != nil {
				//		log.Fatalln("строка 541 ", err)
				//	}
				//	wg.Done()
				//}()
				go func() {
					_, err = q.Create_together(context.Background(), arg)
					if err != nil {
						log.Fatalln("строка 548 ", err)
					}
					wg.Done()
				}()
				wg.Done()
			}()
			wg.Done()
		}(i)

	}
	wg.Wait()
	return err
}
