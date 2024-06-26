// Code generated by diho_bytes_generate world.go; DO NOT EDIT.

package packets

import (
	"context"
	"encoding/binary"
	utils "github.com/Nyarum/diho_bytes_generate/utils"
	"io"
)

func (p *CharacterBoat) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	if err = (&p.CharacterBase).Decode(ctx, reader, endian); err != nil {
		return err
	}
	if err = (&p.CharacterAttribute).Decode(ctx, reader, endian); err != nil {
		return err
	}
	if err = (&p.CharacterKitbag).Decode(ctx, reader, endian); err != nil {
		return err
	}
	if err = (&p.CharacterSkillState).Decode(ctx, reader, endian); err != nil {
		return err
	}
	return nil
}
func (p *Shortcut) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.Type)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.GridID)
	if err != nil {
		return err
	}
	return nil
}
func (p *CharacterShortcut) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	for k := range p.Shortcuts {
		if err = (&p.Shortcuts[k]).Decode(ctx, reader, endian); err != nil {
			return err
		}
	}
	return nil
}
func (p *KitbagItem) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.GridID)
	if err != nil {
		return err
	}
	if p.Filter(ctx, "GridID") == true {
		return err
	}
	err = binary.Read(reader, endian, &p.ID)
	if err != nil {
		return err
	}
	if p.Filter(ctx, "ID") == true {
		return err
	}
	err = binary.Read(reader, endian, &p.Num)
	if err != nil {
		return err
	}
	if p.Filter(ctx, "Num") == true {
		return err
	}
	for k := range p.Endure {
		var tempValue uint16
		if err = binary.Read(reader, endian, &tempValue); err != nil {
			return err
		}
		p.Endure[k] = tempValue
	}
	if p.Filter(ctx, "Endure") == true {
		return err
	}
	for k := range p.Energy {
		var tempValue uint16
		if err = binary.Read(reader, endian, &tempValue); err != nil {
			return err
		}
		p.Energy[k] = tempValue
	}
	if p.Filter(ctx, "Energy") == true {
		return err
	}
	err = binary.Read(reader, endian, &p.ForgeLevel)
	if err != nil {
		return err
	}
	if p.Filter(ctx, "ForgeLevel") == true {
		return err
	}
	err = binary.Read(reader, endian, &p.IsValid)
	if err != nil {
		return err
	}
	if p.Filter(ctx, "IsValid") == true {
		return err
	}
	if p.ID == 3988 {
		err = binary.Read(reader, endian, &p.ItemDBInstID)
		if err != nil {
			return err
		}
	}
	if p.Filter(ctx, "ItemDBInstID") == true {
		return err
	}
	err = binary.Read(reader, endian, &p.ItemDBForge)
	if err != nil {
		return err
	}
	if p.Filter(ctx, "ItemDBForge") == true {
		return err
	}
	if p.ID == 3988 {
		err = binary.Read(reader, endian, &p.BoatNull)
		if err != nil {
			return err
		}
	}
	if p.Filter(ctx, "BoatNull") == true {
		return err
	}
	if p.ID != 3988 {
		err = binary.Read(reader, endian, &p.ItemDBInstID2)
		if err != nil {
			return err
		}
	}
	if p.Filter(ctx, "ItemDBInstID2") == true {
		return err
	}
	err = binary.Read(reader, endian, &p.IsParams)
	if err != nil {
		return err
	}
	if p.Filter(ctx, "IsParams") == true {
		return err
	}
	for k := range p.InstAttrs {
		if err = (&p.InstAttrs[k]).Decode(ctx, reader, endian); err != nil {
			return err
		}
	}
	if p.Filter(ctx, "InstAttrs") == true {
		return err
	}
	return nil
}
func (p *CharacterKitbag) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.Type)
	if err != nil {
		return err
	}
	if p.Type == SYN_KITBAG_INIT {
		err = binary.Read(reader, endian, &p.KeybagNum)
		if err != nil {
			return err
		}
	}
	for k := range p.Items {
		if err = (&p.Items[k]).Decode(ctx, reader, endian); err != nil {
			return err
		}
	}
	return nil
}
func (p *Attribute) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.ID)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.Value)
	if err != nil {
		return err
	}
	return nil
}
func (p *CharacterAttribute) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.Type)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.Num)
	if err != nil {
		return err
	}
	for k := range p.Attributes {
		if err = (&p.Attributes[k]).Decode(ctx, reader, endian); err != nil {
			return err
		}
	}
	return nil
}
func (p *SkillState) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.ID)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.Level)
	if err != nil {
		return err
	}
	return nil
}
func (p *CharacterSkillState) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.StatesLen)
	if err != nil {
		return err
	}
	for k := range p.States {
		if err = (&p.States[k]).Decode(ctx, reader, endian); err != nil {
			return err
		}
	}
	return nil
}
func (p *CharacterSkill) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.ID)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.State)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.Level)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.UseSP)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.UseEndure)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.UseEnergy)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.ResumeTime)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.RangeType)
	if err != nil {
		return err
	}
	for k := range p.Params {
		var tempValue uint16
		if err = binary.Read(reader, endian, &tempValue); err != nil {
			return err
		}
		p.Params[k] = tempValue
	}
	return nil
}
func (p *CharacterSkillBag) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.SkillID)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.Type)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.SkillNum)
	if err != nil {
		return err
	}
	for k := range p.Skills {
		if err = (&p.Skills[k]).Decode(ctx, reader, endian); err != nil {
			return err
		}
	}
	return nil
}
func (p *CharacterAppendLook) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.LookID)
	if err != nil {
		return err
	}
	if p.LookID != 0 {
		err = binary.Read(reader, endian, &p.IsValid)
		if err != nil {
			return err
		}
	}
	return nil
}
func (p *CharacterPK) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.PkCtrl)
	if err != nil {
		return err
	}
	return nil
}
func (p *CharacterLookBoat) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.PosID)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.BoatID)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.Header)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.Body)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.Engine)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.Cannon)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.Equipment)
	if err != nil {
		return err
	}
	return nil
}
func (p *CharacterLookItemSync) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.Endure)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.Energy)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.IsValid)
	if err != nil {
		return err
	}
	return nil
}
func (p *CharacterLookItemShow) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.Num)
	if err != nil {
		return err
	}
	for k := range p.Endure {
		var tempValue uint16
		if err = binary.Read(reader, endian, &tempValue); err != nil {
			return err
		}
		p.Endure[k] = tempValue
	}
	for k := range p.Energy {
		var tempValue uint16
		if err = binary.Read(reader, endian, &tempValue); err != nil {
			return err
		}
		p.Energy[k] = tempValue
	}
	err = binary.Read(reader, endian, &p.ForgeLevel)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.IsValid)
	if err != nil {
		return err
	}
	return nil
}
func (p *CharacterLookItem) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.ID)
	if err != nil {
		return err
	}
	if p.Filter(ctx, "ID") == true {
		return err
	}
	if p.SynType == SynLookChange {
		if err = (&p.ItemSync).Decode(ctx, reader, endian); err != nil {
			return err
		}
	}
	if p.Filter(ctx, "ItemSync") == true {
		return err
	}
	if p.SynType == SynLookSwitch {
		if err = (&p.ItemShow).Decode(ctx, reader, endian); err != nil {
			return err
		}
	}
	if p.Filter(ctx, "ItemShow") == true {
		return err
	}
	err = binary.Read(reader, endian, &p.IsDBParams)
	if err != nil {
		return err
	}
	if p.Filter(ctx, "IsDBParams") == true {
		return err
	}
	for k := range p.DBParams {
		var tempValue uint32
		if err = binary.Read(reader, endian, &tempValue); err != nil {
			return err
		}
		p.DBParams[k] = tempValue
	}
	if p.Filter(ctx, "DBParams") == true {
		return err
	}
	err = binary.Read(reader, endian, &p.IsInstAttrs)
	if err != nil {
		return err
	}
	if p.Filter(ctx, "IsInstAttrs") == true {
		return err
	}
	for k := range p.InstAttrs {
		if err = (&p.InstAttrs[k]).Decode(ctx, reader, endian); err != nil {
			return err
		}
	}
	if p.Filter(ctx, "InstAttrs") == true {
		return err
	}
	return nil
}
func (p *CharacterLookHuman) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.HairID)
	if err != nil {
		return err
	}
	for k := range p.ItemGrid {
		if err = (&p.ItemGrid[k]).Decode(ctx, reader, endian); err != nil {
			return err
		}
	}
	return nil
}
func (p *CharacterLook) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.SynType)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.TypeID)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.IsBoat)
	if err != nil {
		return err
	}
	if p.IsBoat == 1 {
		if err = (&p.LookBoat).Decode(ctx, reader, endian); err != nil {
			return err
		}
	}
	if p.IsBoat == 0 {
		if err = (&p.LookHuman).Decode(ctx, reader, endian); err != nil {
			return err
		}
	}
	return nil
}
func (p *EntityEvent) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.EntityID)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.EntityType)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.EventID)
	if err != nil {
		return err
	}
	p.EventName, err = utils.ReadStringNull(reader)
	if err != nil {
		return err
	}
	return nil
}
func (p *CharacterSide) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.SideID)
	if err != nil {
		return err
	}
	return nil
}
func (p *Position) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.X)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.Y)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.Radius)
	if err != nil {
		return err
	}
	return nil
}
func (p *CharacterBase) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.ChaID)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.WorldID)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.CommID)
	if err != nil {
		return err
	}
	p.CommName, err = utils.ReadStringNull(reader)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.GmLvl)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.Handle)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.CtrlType)
	if err != nil {
		return err
	}
	p.Name, err = utils.ReadStringNull(reader)
	if err != nil {
		return err
	}
	p.MottoName, err = utils.ReadStringNull(reader)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.Icon)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.GuildID)
	if err != nil {
		return err
	}
	p.GuildName, err = utils.ReadStringNull(reader)
	if err != nil {
		return err
	}
	p.GuildMotto, err = utils.ReadStringNull(reader)
	if err != nil {
		return err
	}
	p.StallName, err = utils.ReadStringNull(reader)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.State)
	if err != nil {
		return err
	}
	if err = (&p.Position).Decode(ctx, reader, endian); err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.Angle)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.TeamLeaderID)
	if err != nil {
		return err
	}
	if err = (&p.Side).Decode(ctx, reader, endian); err != nil {
		return err
	}
	if err = (&p.EntityEvent).Decode(ctx, reader, endian); err != nil {
		return err
	}
	if err = (&p.Look).Decode(ctx, reader, endian); err != nil {
		return err
	}
	if err = (&p.PkCtrl).Decode(ctx, reader, endian); err != nil {
		return err
	}
	for k := range p.LookAppend {
		if err = (&p.LookAppend[k]).Decode(ctx, reader, endian); err != nil {
			return err
		}
	}
	return nil
}
func (p *EnterGame) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	err = binary.Read(reader, endian, &p.EnterRet)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.AutoLock)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.KitbagLock)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.EnterType)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.IsNewChar)
	if err != nil {
		return err
	}
	p.MapName, err = utils.ReadStringNull(reader)
	if err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.CanTeam)
	if err != nil {
		return err
	}
	if err = (&p.CharacterBase).Decode(ctx, reader, endian); err != nil {
		return err
	}
	if err = (&p.CharacterSkillBag).Decode(ctx, reader, endian); err != nil {
		return err
	}
	if err = (&p.CharacterSkillState).Decode(ctx, reader, endian); err != nil {
		return err
	}
	if err = (&p.CharacterAttribute).Decode(ctx, reader, endian); err != nil {
		return err
	}
	if err = (&p.CharacterKitbag).Decode(ctx, reader, endian); err != nil {
		return err
	}
	if err = (&p.CharacterShortcut).Decode(ctx, reader, endian); err != nil {
		return err
	}
	err = binary.Read(reader, endian, &p.BoatLen)
	if err != nil {
		return err
	}
	for k := range p.CharacterBoats {
		if err = (&p.CharacterBoats[k]).Decode(ctx, reader, endian); err != nil {
			return err
		}
	}
	err = binary.Read(reader, endian, &p.ChaMainID)
	if err != nil {
		return err
	}
	return nil
}
func (p *EnterGameRequest) Decode(ctx context.Context, reader io.Reader, endian binary.ByteOrder) error {
	var err error
	p.CharacterName, err = utils.ReadStringNull(reader)
	if err != nil {
		return err
	}
	return nil
}
