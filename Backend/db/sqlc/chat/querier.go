// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
    "context"
)


type ChatQuerier interface {
	Createsocialmedia(ctx context.Context, arg CreatesocialmediaParams) (Socialmedium, error)
	Getsocialmedia(ctx context.Context, arg GetsocialmediaParams) (Socialmedium, error)
	Listsocialmedia(ctx context.Context) ([]Socialmedium, error)
	Updatesocialmedia(ctx context.Context, arg UpdatesocialmediaParams) (Socialmedium, error)
	Deletesocialmedia(ctx context.Context, userID int32) error

    CreateChatTable(ctx context.Context, tablename string) error
    CreateRecord(ctx context.Context, arg CreateRecordParams, tablename string) (Record, error)
    DeleteRecord(ctx context.Context, arg DeleteRecordParams, tablename string) error
    GetLastMsg(ctx context.Context, targetID int32, tablename string) (GetLastMsgRow, error) 
    GetRecord(ctx context.Context, targetID int32, tablename string) ([]Record, error)
    GetTargetID(ctx context.Context,tablename string) ([]int32, error)
    GetChatRow(ctx context.Context, tablename string) (int32, error)
	UpdateRead(ctx context.Context, targetID int32,tablename string) error
    UpdateRecord(ctx context.Context, arg UpdateRecordParams, tablename string) (Record, error)
}


var _ ChatQuerier = (*Queries)(nil)