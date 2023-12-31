
namespace go note

struct BaseResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct Note {
    1:i64 note_id
    2:i64 user_id
    3:string user_name
    4:string user_avatar
    5:string title
    6:string content
    7:i64 create_time
}

struct CreateNoteRequest {
    1:string title
    2:string content
    3:i64 user_id
}

struct CreateNoteResponse {
    1:BaseResp base_resp
}

struct DeleteNoteRequest {
    1:i64 note_id
    2:i64 user_id
}

struct DeleteNoteResponse {
    1:BaseResp base_resp
}

struct UpdateNoteRequest {
    1:i64 note_id
    2:i64 user_id
    3:optional string title
    4:optional string content
}

struct UpdateNoteResponse {
    1:BaseResp base_resp
}

struct MGetNoteRequest {
    1:list<i64> note_ids
}

struct MGetNoteResponse {
    1:list<Note> notes
    2:BaseResp base_resp
}

struct QueryNoteRequest {
    1:i64 user_id
    2:optional string search_key
    3:i64 offset
    4:i64 limit
}

struct QueryNoteResponse {
    1:list<Note> notes
    2:i64 total
    3:BaseResp base_resp
}

service NoteService {
    CreateNoteResponse CreateNote(1:CreateNoteRequest req)
    MGetNoteResponse MGetNote(1:MGetNoteRequest req)
    DeleteNoteResponse DeleteNote(1:DeleteNoteRequest req)
    QueryNoteResponse QueryNote(1:QueryNoteRequest req)
    UpdateNoteResponse UpdateNote(1:UpdateNoteRequest req)
}
