
/// 文件服务
///
/// 文件相关的接口
@entity("Files")
interface Files {
    /// 获取文件详情
    @http.get("v1/files")
    get_files(include_deleted: Bool @1 //< 默认值为false
              limit: String @2 //< 默认值为null
              cursor: String @3 //< 默认值为null
              last_updated_after: Timestamp @4 //< 默认值为null
              ) -> [File]

    /// 获取指定的文件详情
    @http.get("v1/files/{id}")
    get_file(id: String @1) -> File

    /// 新增文件
    @http.post("v1/files")
    create_file(file: File @1) -> File

    /// 更新文件信息
    @http.put("v1/files")
    update_file(file: File @1) -> File

    /// 删除用户
    @http.delete("v1/files/{id}")
    delete_file(id: String @1)
}
