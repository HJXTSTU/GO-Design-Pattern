package flyweight

type IProperty interface{
	GetHPLimit()int
	GetMPLimit()int
}

type PeopleBase struct{
	MAX_HP int
	MAX_MP int
}

func NewPeopleBase()*PeopleBase{
	return &PeopleBase{100,100}
}

func (this *PeopleBase)GetHPLimit()int{
	return this.MAX_HP
}

func (this *PeopleBase)GetMPLimit()int{
	return this.MAX_MP
}

type Helmet struct{
	base IProperty
	HP_ADD int
	MP_ADD int
}

func (this *Helmet)GetHPLimit()int{
	return this.base.GetHPLimit()+this.HP_ADD
}

func (this *Helmet)GetMPLimit()int{
	return this.base.GetMPLimit()+this.MP_ADD
}

func NewHelmet(property IProperty,hp_add,mp_add int)*Helmet{
	return &Helmet{property,hp_add,mp_add}
}
