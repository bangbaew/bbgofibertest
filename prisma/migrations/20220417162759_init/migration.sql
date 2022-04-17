/*
  Warnings:

  - You are about to drop the `user` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropTable
PRAGMA foreign_keys=off;
DROP TABLE "user";
PRAGMA foreign_keys=on;

-- CreateTable
CREATE TABLE "User" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "Firstname" TEXT NOT NULL,
    "Lastname" TEXT NOT NULL,
    "PicFormat" TEXT,
    "Address" TEXT NOT NULL,
    "Tel" TEXT NOT NULL,
    "Balance" INTEGER NOT NULL,
    "Bio" TEXT,
    "Email" TEXT NOT NULL,
    "Password" TEXT NOT NULL,
    "Desc" TEXT
);

-- CreateTable
CREATE TABLE "Artisan" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "Firstname" TEXT NOT NULL,
    "Lastname" TEXT NOT NULL,
    "PicFormat" TEXT,
    "Tel" TEXT,
    "ClusterID" TEXT NOT NULL,
    CONSTRAINT "Artisan_ClusterID_fkey" FOREIGN KEY ("ClusterID") REFERENCES "Cluster" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "Cluster" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "Name" TEXT NOT NULL,
    "Region" TEXT NOT NULL,
    "Grade" TEXT NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "User_Email_key" ON "User"("Email");
