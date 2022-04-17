/*
  Warnings:

  - You are about to drop the column `Desc` on the `User` table. All the data in the column will be lost.

*/
-- RedefineTables
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_User" (
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
    "Password" TEXT NOT NULL
);
INSERT INTO "new_User" ("Address", "Balance", "Bio", "Email", "Firstname", "Lastname", "Password", "PicFormat", "Tel", "createdAt", "id", "updatedAt") SELECT "Address", "Balance", "Bio", "Email", "Firstname", "Lastname", "Password", "PicFormat", "Tel", "createdAt", "id", "updatedAt" FROM "User";
DROP TABLE "User";
ALTER TABLE "new_User" RENAME TO "User";
CREATE UNIQUE INDEX "User_Email_key" ON "User"("Email");
PRAGMA foreign_key_check;
PRAGMA foreign_keys=ON;
