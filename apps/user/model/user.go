package model

type User struct {

}

func (u *User) TableName() string {
	return "user"
}

/*
# class WMSOutPermission(models.Model):
#     id = models.AutoField(primary_key=True, auto_created=True)
#
#     class Meta:
#         db_table = "wms_out_permission"
#         managed = False
#
#
# class WMSOutGroup(models.Model):
#     id = models.AutoField(primary_key=True, auto_created=True)
#     warehouse = models.ForeignKey(WMSOutWarehouse, related_name="warehouse_groups", on_delete=models.DO_NOTHING, verbose_name="所属仓库")
#     group_name = models.CharField(max_length=20, unique=True, verbose_name="组名")
#     permissions = models.ManyToManyField(WMSOutPermission, related_name="permission_groups", through="WMSOutGroupPermission", verbose_name="组权限")
#     users = models.ManyToManyField(User, related_name="user_groups", through="WMSOutGroupUser", verbose_name="组员")
#     is_usable = models.BooleanField(default=True, verbose_name="是否可用")
#     creator = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.DO_NOTHING, verbose_name="创建人")
#     create_time = models.DateTimeField(auto_now_add=True, verbose_name="记录创建时间")
#     update_time = models.DateTimeField(auto_now=True, verbose_name="记录更新时间")
#
#     class Meta:
#         db_table = "wms_out_group"
#         managed = False
#
#
# class WMSOutGroupPermission(models.Model):
#     id = models.AutoField(primary_key=True, auto_created=True)
#     group = models.ForeignKey(WMSOutGroup, on_delete=models.DO_NOTHING, verbose_name="组")
#     permission = models.ForeignKey(WMSOutPermission, on_delete=models.DO_NOTHING, verbose_name="权限")
#
#     class Meta:
#         db_table = "wms_out_group_permission"
#         managed = False
#
#
# class WMSOutGroupUser(models.Model):
#     id = models.AutoField(primary_key=True, auto_created=True)
#     group = models.ForeignKey(WMSOutGroup, on_delete=models.DO_NOTHING, verbose_name="组")
#     user = models.ForeignKey(User, on_delete=models.DO_NOTHING, verbose_name="员工")
#
#     class Meta:
#         db_table = "wms_out_group_user"
#         managed = False
*/
