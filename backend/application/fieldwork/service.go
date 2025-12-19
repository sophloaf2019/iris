package fieldwork

import (
	"errors"
	"iris/domain/types/auth"
	"iris/domain/types/fieldwork"
)

// PostDebrief
// EditDebrief
// DeleteDebrief
// PostGOI
// EditGOI
// DeleteGOI
// PostMission
// EditMission
// DeleteMission

type Service struct {
	dbrfRepo DebriefRepo
	goiRepo  GOIRepo
	msnRepo  MissionRepo
}

func NewService(dbrfRepo DebriefRepo, goiRepo GOIRepo, msnRepo MissionRepo) *Service {
	return &Service{dbrfRepo, goiRepo, msnRepo}
}

func (s *Service) GetDebrief(ctx auth.Context, id int) (*fieldwork.Debrief, error) {
	return s.dbrfRepo.Get(id)
}

func (s *Service) GetDebriefs(ctx auth.Context) ([]*fieldwork.Debrief, error) {
	return s.dbrfRepo.GetMany()
}

func (s *Service) GetDebriefBySlug(ctx auth.Context, slug string) (*fieldwork.Debrief, error) {
	return s.dbrfRepo.Slug(slug)
}

func (s *Service) PostDebrief(ctx auth.Context, dbrf *fieldwork.Debrief) error {
	if !auth.Can(
		ctx.Clearance,
		auth.ActionCreate,
		auth.ContentDebrief,
		ctx.UserID == dbrf.AuthorID,
	) {
		return errors.New("access denied")
	}
	return s.dbrfRepo.Create(dbrf)
}

func (s *Service) UpdateDebrief(ctx auth.Context, dbrf *fieldwork.Debrief) error {
	_, err := s.dbrfRepo.Get(dbrf.ID)
	if err != nil {
		return err
	}
	if !auth.Can(
		ctx.Clearance,
		auth.ActionEdit,
		auth.ContentDebrief,
		ctx.UserID == dbrf.AuthorID,
	) {
		return errors.New("access denied")
	}
	return s.dbrfRepo.Update(dbrf)
}

func (s *Service) DeleteDebrief(ctx auth.Context, id int) error {
	dbrf, err := s.dbrfRepo.Get(id)
	if err != nil {
		return err
	}
	if !auth.Can(
		ctx.Clearance,
		auth.ActionEdit,
		auth.ContentDebrief,
		ctx.UserID == dbrf.AuthorID,
	) {
		return errors.New("access denied")
	}
	return s.dbrfRepo.Delete(dbrf.ID)
}

func (s *Service) GetGOI(ctx auth.Context, id int) (*fieldwork.GOI, error) {
	return s.goiRepo.Get(id)
}

func (s *Service) GetGOIs(ctx auth.Context) ([]*fieldwork.GOI, error) {
	return s.goiRepo.GetMany()
}

func (s *Service) GetGOIBySlug(ctx auth.Context, slug string) (*fieldwork.GOI, error) {
	return s.goiRepo.Slug(slug)
}

func (s *Service) PostGOI(ctx auth.Context, goi *fieldwork.GOI) error {
	if !auth.Can(
		ctx.Clearance,
		auth.ActionCreate,
		auth.ContentGOI,
		ctx.UserID == goi.AuthorID,
	) {
		return errors.New("access denied")
	}
	return s.goiRepo.Create(goi)
}

func (s *Service) UpdateGOI(ctx auth.Context, goi *fieldwork.GOI) error {
	_, err := s.goiRepo.Get(goi.ID)
	if err != nil {
		return err
	}
	if !auth.Can(
		ctx.Clearance,
		auth.ActionEdit,
		auth.ContentGOI,
		ctx.UserID == goi.AuthorID,
	) {
		return errors.New("access denied")
	}
	return s.goiRepo.Update(goi)
}

func (s *Service) DeleteGOI(ctx auth.Context, id int) error {
	goi, err := s.goiRepo.Get(id)
	if err != nil {
		return err
	}
	if !auth.Can(
		ctx.Clearance,
		auth.ActionEdit,
		auth.ContentGOI,
		ctx.UserID == goi.AuthorID,
	) {
		return errors.New("access denied")
	}
	return s.goiRepo.Delete(goi.ID)
}

func (s *Service) GetMission(ctx auth.Context, id int) (*fieldwork.Mission, error) {
	return s.msnRepo.Get(id)
}

func (s *Service) GetMissions(ctx auth.Context) ([]*fieldwork.Mission, error) {
	return s.msnRepo.GetMany()
}

func (s *Service) GetMissionBySlug(ctx auth.Context, slug string) (*fieldwork.Mission, error) {
	return s.msnRepo.Slug(slug)
}

func (s *Service) PostMission(ctx auth.Context, msn *fieldwork.Mission) error {
	if !auth.Can(
		ctx.Clearance,
		auth.ActionCreate,
		auth.ContentMission,
		ctx.UserID == msn.AuthorID,
	) {
		return errors.New("access denied")
	}
	return s.msnRepo.Create(msn)
}

func (s *Service) UpdateMission(ctx auth.Context, msn *fieldwork.Mission) error {
	_, err := s.msnRepo.Get(msn.ID)
	if err != nil {
		return err
	}
	if !auth.Can(
		ctx.Clearance,
		auth.ActionEdit,
		auth.ContentMission,
		ctx.UserID == msn.AuthorID,
	) {
		return errors.New("access denied")
	}
	return s.msnRepo.Update(msn)
}

func (s *Service) DeleteMission(ctx auth.Context, id int) error {
	msn, err := s.msnRepo.Get(id)
	if err != nil {
		return err
	}
	if !auth.Can(
		ctx.Clearance,
		auth.ActionEdit,
		auth.ContentMission,
		ctx.UserID == msn.AuthorID,
	) {
		return errors.New("access denied")
	}
	return s.msnRepo.Delete(msn.ID)
}
