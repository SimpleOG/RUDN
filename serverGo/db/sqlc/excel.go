package db

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgtype"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
	"sync"
)

const (
	//тупо счётчик
	//Образовательная программа
	//educationForma = iota // Форма обучения
	//lvlOp                // Уровень ОП
	a = iota
	b
	c // Шифр ООП РУДН
	d // Код направления
	e // Наименование программы
	// Дисциплина или вид учебной  работы
	f
	g
	h
	i
	j
	k
	l
	m
	n
	o
	p
	q
	r
	s
	t
	u
	v
	//контингент обучающихся
	w
	x
	y
	z
	aa
	ab
	ac
	ad
	ae
	af
	//Сведения о ППС
	ag
	ah
	ai
	aj
	ak
	//Объём учебной работы ППС
	al
	am
	an
	ao
	ap
	aq
	ar
	as
	at
	au
	av
	aw
	ax
	ay
	az
	ba
	bb
	bc
	bd
	be
	bf
	bg
	bh
	bi
	bj
	bk
	bl
	start = 5
	end   = 1640
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

func (qur *Queries) ReadEducationalProgram(data [][]string) ([size]EducationalProgram, error) {
	lock := new(sync.Mutex)
	ma := make(map[int]EducationalProgram)
	edu_form := data[a]
	level_op := data[b]
	theCodeOfTheOOPRUDN := data[c]
	directionCode := data[d]
	nameOfTheProgram := data[e]
	var wg sync.WaitGroup
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {
			arg := Create_EducationalProgramParams{
				edu_form[i],
				level_op[i],
				theCodeOfTheOOPRUDN[i],
				directionCode[i],
				nameOfTheProgram[i],
			}
			if arg.NameOfTheProgram == "" || arg.DirectionCode == "" || arg.TheCodeOfTheOopRudn == "" {
				fmt.Println(i)
			}
			n, err := qur.Create_EducationalProgram(context.Background(), arg)
			if err != nil {
				log.Fatalln(err)
			}
			lock.Lock()
			ma[i] = n
			lock.Unlock()
			wg.Done()
		}(i)

	}
	wg.Wait()
	var arr [size]EducationalProgram
	for i, v := range ma {
		arr[i] = v
	}
	return arr, nil
}

func (qur *Queries) ReadDisciplineOrTypeOfAcademicWork(data [][]string) ([size]DisciplineOrTypeOfAcademicWork, error) {
	var err error
	lock := new(sync.Mutex)
	ma := make(map[int]DisciplineOrTypeOfAcademicWork)
	Block := data[f]
	Component := data[g]
	Nvrup := data[h]
	Nameof := data[j]
	Dopinfo := data[i]
	var wg sync.WaitGroup
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func(i int) {
			arg := Create_Discipline_or_type_of_academic_workParams{
				Block:                                   Block[i],
				Component:                               Component[i],
				NVRup:                                   Nvrup[i],
				NameOfTheDisciplineOrTypeOfAcademicWork: Nameof[i],
				DopInfo:                                 Dopinfo[i],
			}
			c, err := qur.Create_Discipline_or_type_of_academic_work(context.Background(), arg)
			if err != nil {
				log.Fatalln(err)
			}
			lock.Lock()
			ma[i] = c
			lock.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	var arr [size]DisciplineOrTypeOfAcademicWork
	for i, v := range ma {
		arr[i] = v
	}
	return arr, err
}

func (qur *Queries) ReadKW(data [][]string) ([size]KW, error) {
	var err error
	lock := new(sync.Mutex)
	ma := make(map[int]KW)
	SemesterOrModule := data[k]
	WeeksPerSemesterModule := data[l]
	TypeOfEducationalWork := data[m]
	LectureHours := data[n]
	LaboratoriesHours := data[o]
	PractiseHours := data[p]
	TypeOfPaOrGia := data[q]
	CourseWorks := data[r]
	CourseProjects := data[s]
	CourseUchAveZeOnRup := data[t]
	PrZeOnRup := data[u]
	nirZeByRup := data[v]
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
				TypeOfPaOrGia:          TypeOfPaOrGia[i],
				CourseWorks:            CourseWorks[i],
				CourseProjects:         CourseProjects[i],
				CourseUchAveZeOnRup:    CourseUchAveZeOnRup[i],
				PrZeOnRup:              PrZeOnRup[i],
				NirZeByRup:             nirZeByRup[i],
			}
			c, err := qur.Create_k_w(context.Background(), arg)
			lock.Lock()
			ma[i] = c
			lock.Unlock()
			if err != nil {
				log.Fatalln(err)
			}
			wg.Done()
		}(i)

	}
	wg.Wait()
	var arr [size]KW
	for i, v := range ma {
		arr[i] = v
	}
	return arr, err
}

func (qur *Queries) ReadTheContingentOfStudents(data [][]string) ([size]TheContingentOfStudent, error) {
	var err error
	lock := new(sync.Mutex)
	ma := make(map[int]TheContingentOfStudent)
	Code := data[w]
	groupNumber := data[x]
	ofGroups := data[y]
	Subgroups := data[z]
	totalPeople := data[aa]
	rf := data[ab]
	Foreign := data[ac]
	standart := data[ad]
	Calculated := data[ae]
	pk := data[af]
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
				Rf:          rf[i],
				Foreign:     Foreign[i],
				Standard:    standart[i],
				Calculated:  Calculated[i],
				Pk:          pk[i],
			}
			c, err := qur.Create_the_contingent_of_students(context.Background(), arg)
			if err != nil {
				log.Fatalln(err)
			}
			lock.Lock()
			ma[i] = c
			lock.Unlock()
			wg.Done()
		}(i)

	}
	wg.Wait()
	var arr [size]TheContingentOfStudent
	for i, v := range ma {
		arr[i] = v
	}
	return arr, err
}

func (qur *Queries) ReadInformationAboutPps(data [][]string) ([size]InformationAboutPP, error) {
	var err error
	lock := new(sync.Mutex)
	ma := make(map[int]InformationAboutPP)
	Department := data[ag]
	Post := data[ah]
	TermsOfAttraction := data[ai]
	FullName := data[aj]
	ASpecialFeature := data[ak]
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
			c, err := qur.Create_information_about_PPS(context.Background(), arg)
			if err != nil {
				log.Fatalln(err)
			}
			lock.Lock()
			ma[i] = c
			lock.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	var arr [size]InformationAboutPP
	for i, v := range ma {
		arr[i] = v
	}
	return arr, err
}

func (qur *Queries) ReadTheAmountOfTeachingWorkOfTheTeachingStaff(data [][]string) ([size]TheAmountOfTeachingWorkOfTheTeachingStaff, error) {
	var err error
	lock := new(sync.Mutex)
	ma := make(map[int]TheAmountOfTeachingWorkOfTheTeachingStaff)
	Lectures := data[al]
	PracticeOrSeminars := data[am]
	LabWorksOrClinicalClasses := data[an]
	CurrentControl := data[ao]
	InterimCertificationPoForBrs := data[ap]
	RegistrationOfPaResults := data[aq]
	OngoingConsultationsOnTheDiscipline := data[ar]
	CourseWorks := data[as]
	CourseProjects := data[at]
	EducationalPractice := data[au]
	ProcPedagogicalAndPreGraduatePractices := data[av]
	nir := data[aw]
	PracticesIncludingResearchOfDigitalMagistracies := data[ax]
	ReviewingTheAbstractsOfGraduateStudents := data[ay]
	CandidatesExama := data[az]
	ScientificGuidance := data[ba]
	TheLeadershipOfTheWrcOrTheNkr := data[bb]
	ReviewOfTheWrc := data[bc]
	gek := data[bd]
	Total := data[be]
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
			cand, err := strconv.ParseFloat(CandidatesExama[i], 10)
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
				InterimCertificationPoForBrs:           inter,
				RegistrationOfPaResults:                reg,
				OngoingConsultationsOnTheDiscipline:    ong,
				CourseWorks:                            course,
				CourseProjects:                         projects,
				EducationalPractice:                    edu,
				ProcPedagogicalAndPreGraduatePractices: proc,
				Nir:                                    n,
				PracticesIncludingResearchOfDigitalMagistracies: prac,
				ReviewingTheAbstractsOfGraduateStudents:         rev,
				CandidatesExam:                                  cand,
				ScientificGuidance:                              sci,
				TheLeadershipOfTheWrcOrTheNkr:                   leader,
				ReviewOfTheWrc:                                  wrc,
				Gek:                                             ge,
				Total:                                           to,
			}

			amount, err := qur.Create_the_amount_of_teaching_work_of_the_teaching_staff(context.Background(), arg)

			if err != nil {
				log.Fatalln(err)
			}
			lock.Lock()
			ma[i] = amount
			lock.Unlock()
			wg.Done()
		}(i)

	}
	wg.Wait()
	var arr [size]TheAmountOfTeachingWorkOfTheTeachingStaff
	for i, v := range ma {
		arr[i] = v
	}
	return arr, err
}

//func (qur *Queries) ReadSemesters(data [][]string) ([size]Semester, error) {
//	lock := new(sync.Mutex)
//	ma := make(map[int]Semester)
//	if data[bg][0]!=""{
//
//	}
//	AuditionWork:= data[c]
//	PairsPerWeek := data[d]
//	activities := data[e]
//	var wg sync.WaitGroup
//	for i := 0; i < size; i++ {
//		wg.Add(1)
//		go func(i int) {
//			arg := Create_SemesterParams{
//				[i],
//				directionCode[i],
//				nameOfTheProgram[i],
//			}
//			if arg.NameOfTheProgram == "" || arg.DirectionCode == "" || arg.TheCodeOfTheOOPRUDN == "" {
//				fmt.Println(i)
//			}
//			n, err := qur.Create_EducationalProgram(context.Background(), arg)
//			if err != nil {
//				log.Fatalln(err)
//			}
//			lock.Lock()
//			ma[i] = n
//			lock.Unlock()
//			wg.Done()
//		}(i)
//
//	}
//	wg.Wait()
//	var arr [size]EducationalProgram
//	for i, v := range ma {
//		arr[i] = v
//	}
//	return arr, nil
//}

func (qur *Queries) ReadItAll() error {
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
			program, err = qur.ReadEducationalProgram(data)
			if err != nil {
				log.Fatalln(err)
			}
			wg.Done()
		}()
		go func() {
			discipline, err = qur.ReadDisciplineOrTypeOfAcademicWork(data)
			if err != nil {
				log.Fatalln(err)
			}
			wg.Done()
		}()
		go func() {
			kw, err = qur.ReadKW(data)
			if err != nil {
				log.Fatalln(err)
			}
			wg.Done()
		}()
		go func() {
			group, err = qur.ReadTheContingentOfStudents(data)
			if err != nil {
				log.Fatalln(err)
			}
			wg.Done()
		}()
		go func() {
			pps, err = qur.ReadInformationAboutPps(data)
			if err != nil {
				log.Fatalln(err)
			}
			wg.Done()
		}()
		go func() {
			amount, err = qur.ReadTheAmountOfTeachingWorkOfTheTeachingStaff(data)
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
				ProgramID:    program[i].ID,
				DisciplineID: discipline[i].ID,
				GroupID:      group[i].ID,
				TeacherID:    pps[i].ID,
				KWID:         kw[i].ID,
				AmountID:     amount[i].ID,
			}
			wg.Add(2)
			go func() {

				go func() {
					_, err = qur.Create_together(context.Background(), arg)
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
func (qur *Queries) TakeInfo(fields []string, name string) error {
	field := ""
	for i, v := range fields {
		if i == len(fields)-1 {
			field += fmt.Sprintf("'%s' ", v)
			continue
		}
		field += fmt.Sprintf(" '%s' ,", v)
	}
	query := `SELECT ` + field + ` 
	from discipline_or_type_of_academic_work d
	join together t on d.id = t.discipline_id 
    join  k_w kw on t.k_w_id = kw.id 	
	join the_amount_of_teaching_work_of_the_teaching_staff taotwotts on t.amount_id = taotwotts.id 
	join "information_about_PPS" iaP on iaP.id = t.teacher_id and iap.full_name= $1
	join the_contingent_of_students tcos on t.group_id = tcos.id
	join educational_program ep on t.program_id = ep.id`
	rows, err := qur.db.Query(context.Background(), query, name)
	if err != nil {
		return err
	}
	fmt.Println(rows)
	return nil
}
