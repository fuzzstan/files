
type File {
    id: String @1

    name: String @2

    size: Int @10

    permalink: String @15

    date_added: Timestamp @100

    last_updated: Timestamp @101

    is_deleted: Bool @102

//    /// create time
//    create_time: Timestamp @100
//
//    /// update time
//    update_time: Timestamp @101
//
//    /// delete time
//    delete_time: Timestamp @102
}
