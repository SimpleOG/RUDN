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
	end   = 1600
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
func fill(m map[int]int32) [size]int32 {
	var arr [size]int32
	for i, v := range m {
		arr[i] = v
	}
	return arr
}

func (q *Queries) ReadEducationalProgram(data [][]string) ([size]int32, error) {
	lock := new(sync.Mutex)
	m := make(map[int]int32)
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
			m[i] = n.ID
			lock.Unlock()
			wg.Done()
		}(i)

	}
	wg.Wait()
	arr := fill(m)
	return arr, nil
}

func (q *Queries) ReadDisciplineOrTypeOfAcademicWork(data [][]string) ([size]int32, error) {
	var err error
	lock := new(sync.Mutex)
	m := make(map[int]int32)
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
			m[i] = c.ID
			lock.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	arr := fill(m)
	return arr, err
}

func (q *Queries) ReadKW(data [][]string) ([size]int32, error) {
	var err error
	lock := new(sync.Mutex)
	m := make(map[int]int32)
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
			m[i] = c.ID
			lock.Unlock()
			if err != nil {
				log.Fatalln(err)
			}
			wg.Done()
		}(i)

	}
	wg.Wait()
	arr := fill(m)
	return arr, err
}

func (q *Queries) ReadTheContingentOfStudents(data [][]string) ([size]int32, error) {
	var err error
	lock := new(sync.Mutex)
	m := make(map[int]int32)
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
			m[i] = c.ID
			lock.Unlock()
			wg.Done()
		}(i)

	}
	wg.Wait()
	arr := fill(m)
	return arr, err
}

func (q *Queries) ReadInformationAboutPps(data [][]string) ([size]int32, error) {
	var err error
	lock := new(sync.Mutex)
	m := make(map[int]int32)
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
			m[i] = c.ID
			lock.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	arr := fill(m)
	return arr, err
}

func (q *Queries) ReadTheAmountOfTeachingWorkOfTheTeachingStaff(data [][]string) ([size]int32, error) {
	var err error
	lock := new(sync.Mutex)
	m := make(map[int]int32)
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

			c, err := q.Create_the_amount_of_teaching_work_of_the_teaching_staff(context.Background(), arg)

			if err != nil {
				log.Fatalln(err)
			}
			lock.Lock()
			m[i] = c.ID
			lock.Unlock()
			wg.Done()
		}(i)

	}
	wg.Wait()
	arr := fill(m)
	return arr, err
}

func (q *Queries) ReadItAll() error {
	var err error
	data, err := ReadExcel()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Всего прочитано элементов ", len(data)*end)
	wg := new(sync.WaitGroup)
	var program [size]int32
	var discipline [size]int32
	var kw [size]int32
	var group [size]int32
	var pps [size]int32
	var amount [size]int32
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
				ProgramID:    program[i],
				DisciplineID: discipline[i],
				GroupID:      group[i],
				TeacherID:    pps[i],
				KWID:         kw[i],
				AmountID:     amount[i],
			}
			_, err = q.Create_together(context.Background(), arg)
			wg.Done()
		}(i)
	}
	wg.Wait()
	return err
}
