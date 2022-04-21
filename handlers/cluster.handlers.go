package handlers

import (
	prisma "boilerplate/database"
	"boilerplate/models"
	"boilerplate/prisma/db"

	"github.com/gofiber/fiber/v2"
)

type ClusterHandler struct{}

// @Summary Show all users
// @Description Show all users in the database
// @Tags Cluster
// Param id path int true "Account ID"
// @Router /api/v1/clusters [get]
func (ClusterHandler) FindMany(c *fiber.Ctx) error {

	result, err := prisma.Client.Cluster.FindMany().Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	PrintJSON(result)

	return c.JSON(result)
}

// @Tags Cluster
// @Param id path string true "Account ID"
// @Router /api/v1/clusters/{id} [get]
func (ClusterHandler) Find(c *fiber.Ctx) error {

	result, err := prisma.Client.Cluster.FindUnique(
		db.Cluster.ID.Equals(c.Params("id")),
	).Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	PrintJSON(result)

	return c.JSON(result)
}

// @Tags Cluster
// UserCreate registers a user
//@Param        data  body      models.Cluster  true  "Account Info"
// @Router /api/v1/clusters [post]
func (ClusterHandler) Create(c *fiber.Ctx) error {

	var payload models.Cluster

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	PrintJSON(payload)

	created, err := prisma.Client.Cluster.CreateOne(
		db.Cluster.Name.Set(payload.Name),
		db.Cluster.Region.Set(payload.Region),
		db.Cluster.Grade.Set(payload.Grade),
	).Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	PrintJSON(created)

	return c.JSON(created)
}

// @Tags Cluster
// @Param id path string true "Account ID"
// @Router /api/v1/clusters/{id} [delete]
func (ClusterHandler) Delete(c *fiber.Ctx) error {
	deleted, err := prisma.Client.Cluster.FindUnique(db.Cluster.ID.Equals(c.Params("id"))).Delete().Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	return c.JSON(deleted)
}

// @Tags Cluster
// @Param id path string true "Account ID"
//@Param        data  body      models.Cluster  true  "Account Info"
// @Router /api/v1/clusters/{id} [patch]
func (ClusterHandler) Update(c *fiber.Ctx) error {
	var payload models.Cluster

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	updated, err := prisma.Client.Cluster.FindUnique(db.Cluster.ID.Equals(c.Params("id"))).Update(
		db.Cluster.Name.Set(payload.Name),
		db.Cluster.Region.Set(payload.Region),
		db.Cluster.Grade.Set(payload.Grade),
	).Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	return c.JSON(updated)
}

// @Tags Cluster
// @Router /api/v1/clusters [delete]
func (ClusterHandler) DeleteMany(c *fiber.Ctx) error {
	deleted, err := prisma.Client.Cluster.FindMany().Delete().Exec(prisma.Ctx)
	if err != nil {
		return err
	}

	return c.JSON(deleted)
}
