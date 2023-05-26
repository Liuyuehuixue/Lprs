package mysql

import (
	"fmt"
	"lpr_system/models"
)

// GetALLSelectedSurveyItems 查询七个表
func GetALLSelectedSurveyItems(id int, age int, testN string) (data models.SelectedSevenTableItems, err error) {
	data.ID = id
	data.TestN = testN
	data.Age = age
	sqlStr := `select IFNULL(survey_sub_range,"无") as survey_sub_range, survey_project_cnname,survey_project_enname
	from base1_gzj
	where min_age<=? and max_age>?
	`
	data.AllItems[0].TabelName = "儿童感知觉发展能力发展调查表"
	err = db.Select(&data.AllItems[0].Items, sqlStr, age, age)

	sqlStr1 := `select IFNULL(survey_sub_range,"无") as survey_sub_range, survey_project_cnname,survey_project_enname
	from base2_cddz
	where min_age<=? and max_age>?
	`
	data.AllItems[1].TabelName = "儿童粗大动作能力发展调查表"
	err = db.Select(&data.AllItems[1].Items, sqlStr1, age, age)
	fmt.Println(data.AllItems[1].Items)

	sqlStr2 := `select IFNULL(survey_sub_range,"无") as survey_sub_range, survey_project_cnname,survey_project_enname
	from base3_jxdz
	where min_age<=? and max_age>?
	`
	data.AllItems[2].TabelName = "儿童精细动作能力发展调查表"
	err = db.Select(&data.AllItems[2].Items, sqlStr2, age, age)
	fmt.Println(data.AllItems[2].Items)

	sqlStr3 := `select IFNULL(survey_sub_range,"无") as survey_sub_range, survey_project_cnname,survey_project_enname
	from base4_yyygt
	where min_age<=? and max_age>?
	`
	data.AllItems[3].TabelName = "儿童语言与沟通能力发展调查表"
	err = db.Select(&data.AllItems[3].Items, sqlStr3, age, age)

	sqlStr4 := `select IFNULL(survey_sub_range,"无") as survey_sub_range, survey_project_cnname,survey_project_enname
	from base5_rzfz
	where min_age<=? and max_age>?
	`
	data.AllItems[4].TabelName = "儿童认知发展能力发展调查表"
	err = db.Select(&data.AllItems[4].Items, sqlStr4, age, age)

	sqlStr5 := `select IFNULL(survey_sub_range,"无") as survey_sub_range, survey_project_cnname,survey_project_enname
	from base6_shjw
	where min_age<=? and max_age>?
	`
	data.AllItems[5].TabelName = "儿童社会交往能力发展调查表"
	err = db.Select(&data.AllItems[5].Items, sqlStr5, age, age)

	sqlStr6 := `select IFNULL(survey_sub_range,"无") as survey_sub_range, survey_project_cnname,survey_project_enname
	from base7_shzl
	where min_age<=? and max_age>?
	`
	data.AllItems[6].TabelName = "儿童社会交往能力发展调查表"
	err = db.Select(&data.AllItems[6].Items, sqlStr6, age, age)

	return
}

// GetTimesByID 获取某孩子有那几次记录
func GetTimesByID(cid int) (data []string, err error) {
	data = make([]string, 0, 4)
	sqlStr := `select test_n from survey_gzj_info where child_id=?`
	err = db.Select(&data, sqlStr, cid)
	return
}

// GetStatusInfo 获取status:缓存或提交
func GetStatusInfo(cid int, test_n string) (stat int, err error) {

	sqlStr := `select status from survey_gzj_info where child_id=? and test_n=?`
	err = db.Get(&stat, sqlStr, cid, test_n)
	fmt.Println(stat, err)
	if err != nil {
		return 2, err
	}
	return
}

// GetCnNameByIDInT1 获取表一所有项目中文名
func GetCnNameByIDInT1() (data []string, err error) {
	sqlStr := `select survey_project_cnname
                        from base1_gzj`
	err = db.Select(&data, sqlStr)
	return
}

// GetOneInfoInT1 获取表1某孩子某次记录
func GetOneInfoInT1(cid int, test_n string) (data models.SurveyGzj, err error) {

	sqlStr := `select *
               from survey_gzj_info where child_id=? and test_n=?`

	err = db.QueryRowx(sqlStr, cid, test_n).StructScan(&data)
	//fmt.Println(data)
	return
}

// GetOneInfoInT2 获取表1某孩子某次记录
func GetOneInfoInT2(cid int, test_n string) (data models.SurveyCddz, err error) {

	sqlStr := `select *
               from survey_cddz_info where child_id=? and test_n=?`

	err = db.QueryRowx(sqlStr, cid, test_n).StructScan(&data)
	//fmt.Println(data)
	return
}

// GetOneInfoInT3 获取表1某孩子某次记录
func GetOneInfoInT3(cid int, test_n string) (data models.SurveyJxdz, err error) {

	sqlStr := `select *
               from survey_jxdz_info where child_id=? and test_n=?`

	err = db.QueryRowx(sqlStr, cid, test_n).StructScan(&data)
	//fmt.Println(data)
	return
}

// GetOneInfoInT4 获取表1某孩子某次记录
func GetOneInfoInT4(cid int, test_n string) (data models.SurveyYyygt, err error) {

	sqlStr := `select *
               from survey_yyygt_info where child_id=? and test_n=?`

	err = db.QueryRowx(sqlStr, cid, test_n).StructScan(&data)
	//fmt.Println(data)
	return
}

// GetOneInfoInT5 获取表1某孩子某次记录
func GetOneInfoInT5(cid int, test_n string) (data models.SurveyRzfz, err error) {

	sqlStr := `select *
               from survey_rzfz_info where child_id=? and test_n=?`

	err = db.QueryRowx(sqlStr, cid, test_n).StructScan(&data)
	//fmt.Println(data)
	return
}

// GetOneInfoInT6 获取表1某孩子某次记录
func GetOneInfoInT6(cid int, test_n string) (data models.SurveyShjw, err error) {

	sqlStr := `select *
               from survey_shjw_info where child_id=? and test_n=?`

	err = db.QueryRowx(sqlStr, cid, test_n).StructScan(&data)
	//fmt.Println(data)
	return
}

// GetOneInfoInT7 获取表1某孩子某次记录
func GetOneInfoInT7(cid int, test_n string) (data models.SurveyShzl, err error) {

	sqlStr := `select *
               from survey_shzl_info where child_id=? and test_n=?`

	err = db.QueryRowx(sqlStr, cid, test_n).StructScan(&data)
	//fmt.Println(data)
	return
}

func GetMaxIdInfo1() (max_id int, err error) {

	sqlStr := `select max(survey_info_id)
               from survey_gzj_info`

	err = db.Get(&max_id, sqlStr)
	fmt.Println(max_id)
	return
}

func SearchSuggest(age, survey_id int) (sug string, err error) {

	switch survey_id {
	case 1:
		println("case 1")
		println(age)
		sqlStr := `select IFNULL(subject_perceptual_perception,"") as subject_perceptual_perception
               from suggest_base  where min_age<=? and max_age>=?`
		err = db.Get(&sug, sqlStr, age, age)
		println(sug)
	case 2:
		println("case 2")
		println(age)
		sqlStr := `select IFNULL(subect_gross_movements,"") as subect_gross_movements
               from suggest_base  where min_age<=? and max_age>=?`
		err = db.Get(&sug, sqlStr, age, age)
		println(sug)
	case 3:
		println("case 3")
		println(age)
		sqlStr := `select IFNULL(subject_fine_movements,"") as subject_fine_movements
               from suggest_base  where min_age<=? and max_age>=?`
		err = db.Get(&sug, sqlStr, age, age)
		println(sug)
	case 4:
		println("case 4")
		println(age)
		sqlStr := `select IFNULL(subject_Language_skills,"") as subject_Language_skills
               from suggest_base  where min_age<=? and max_age>=?`
		err = db.Get(&sug, sqlStr, age, age)
		println(sug)
	case 5:
		println("case 5")
		println(age)
		sqlStr := `select IFNULL(subject_cognitive_abilities,"") as subject_cognitive_abilities
               from suggest_base  where min_age<=? and max_age>=?`
		err = db.Get(&sug, sqlStr, age, age)
		println(sug)
	case 6:
		println("case 6")
		println(age)
		sqlStr := `select IFNULL(subject_social_skills,"") as subject_social_skills
               from suggest_base  where min_age<=? and max_age>=?`
		err = db.Get(&sug, sqlStr, age, age)
		println(sug)
	case 7:
		println("case 7")
		println(age)
		sqlStr := `select IFNULL(subject_liveby_own,"") as subject_liveby_own
               from suggest_base  where min_age<=? and max_age>=?`
		err = db.Get(&sug, sqlStr, age, age)
		println(sug)
	}
	return
}

func SelectSugID(child_id int, test_n string) (n string, err error) {
	var child int

	sqlStr := "select child_id from suggestion where child_id=? and test_n=?"
	if err := db.Get(&child, sqlStr, child_id, test_n); err != nil {
		return "-1", err
	}
	return
}

func InsertSuggest(child_id int, test_n, sug string) (err error) {

	sqlStr := "insert into suggestion (child_id, test_n, suggestion) values (?,?,?)"
	db.Exec(sqlStr, child_id, test_n, sug)
	return
}

func UpdateSuggest(child_id int, test_n, sug string) (err error) {

	sqlStr := "update suggest set suggestion=? where child_id = ? and test_n=?"
	db.Exec(sqlStr, child_id, test_n, sug)
	return
}

func InsertSurveyInfo(sevenSurvey models.SevenSurvey) (err error) {

	//tx, err := db.Beginx() // 开启事务
	//if err != nil {
	//	fmt.Printf("begin trans failed, err:%v\n", err)
	//	return err
	//}
	//defer func() {
	//	if p := recover(); p != nil {
	//		tx.Rollback()
	//		panic(p) // re-throw panic after Rollback
	//	} else if err != nil {
	//		fmt.Println("rollback")
	//		tx.Rollback() // err is non-nil; don't change it
	//	} else {
	//		err = tx.Commit() // err is nil; if Commit returns error update err
	//		fmt.Println("commit")
	//	}
	//}()

	sqlStr := "INSERT INTO survey_gzj_info (survey_info_id,child_id,age, status, test_n,1_gxcj,2_yscj,3_lhzs,4_pdwt,5_ksyd,6_sbmk,7_bryx,8_brwp,9_brys,10_brys,11_btys,12_txpk,13_dwpk,14_wppb,15_dwpt,16_jmpd,17_bxxw,18_thxq,19_gyy,20_jmfy,21_tffy,22_tzsy,23_zysy,24_ltyc,25_ltmb,26_jmfy,27_ltxb,28_ltsz,29_ltyl,30_fbsy,31_fbsy,32_btyq,33_fbls,34_pzfy,35_cjfy,36_ysfy,37_cjpd,38_cjtx,39_qfqz,40_qfgs,41_qfxz,42_dxpl,43_qzpl,44_wdfy,45_fbst,46_fbxk,47_bblr,48_bbnd,49_bbry,50_bbhh,51_szqw,52_bbxc,53_qwpd,54_scsw,55_scsg,p_n,e_n,f_n,score) VALUES (:survey_info_id, :child_id,:age, :status, :test_n,:1_gxcj,:2_yscj,:3_lhzs,:4_pdwt,:5_ksyd,:6_sbmk,:7_bryx,:8_brwp,:9_brys,:10_brys,:11_btys,:12_txpk,:13_dwpk,:14_wppb,:15_dwpt,:16_jmpd,:17_bxxw,:18_thxq,:19_gyy,:20_jmfy,:21_tffy,:22_tzsy,:23_zysy,:24_ltyc,:25_ltmb,:26_jmfy,:27_ltxb,:28_ltsz,:29_ltyl,:30_fbsy,:31_fbsy,:32_btyq,:33_fbls,:34_pzfy,:35_cjfy,:36_ysfy,:37_cjpd,:38_cjtx,:39_qfqz,:40_qfgs,:41_qfxz,:42_dxpl,:43_qzpl,:44_wdfy,:45_fbst,:46_fbxk,:47_bblr,:48_bbnd,:49_bbry,:50_bbhh,:51_szqw,:52_bbxc,:53_qwpd,:54_scsw,:55_scsg,:p_n,:e_n,:f_n,:score)"
	_, err = db.NamedExec(sqlStr, &sevenSurvey.SurveyGzj)
	fmt.Println("gzj")
	fmt.Println(err)

	sqlStr1 := "INSERT INTO survey_cddz_info (survey_info_id,child_id,age,status,test_n,1_zzssldzdqg,2_zzzyzzzzd,3_zadswysjdswj,4_dlzd5s,5_zlsnwysjdswp,6_djz5s,7_djz10szyll,8_jjz8s,9_lhpx,10_pslt,11_pxlt,12_tbyd,13_zwzq,14_zzzq,15_ydzq,16_zzdzmf,17_wqdszcxqtbfy,18_wcdszcxptbfy,19_whdszcxhtbfy,20_djzyb6syb3s,21_fysz,22_dzxz5s,23_ssbdwjxqz,24_cxxz,25_zzx,26_fwslt,27_fwxlt,28_z15cmkdphmsjtxz,29_lbsyjlt,30_lbxyjlt,31_ybsyjlt,32_ybxyjlt,33_djz10b,34_zpdsdtzzx5b,35_jjtjhgz3m,36_jhgxz3m,37_sftt,38_ydtsjld5cm,39_xqt10cm,40_zzlthtjswxt,41_tsg5cmdtj,42_tggdjxdlg,43_zyt3s,44_xqktyb,45_lxkmyb10b,46_lxchb3m,47_lxxht10c,48_lxdjt8c,49_5snz3mjlnlhp,50_zztqz1,51_dstwqz1,52_zstq10c,53_dsdbbsbx3m,54_ssdcwtpbx3m,55_dssgjxqpq1,56_sssgjxqpq1,57_ssxxpqz1,58_dsxxpqz1,59_dssgjpzkdz3my2mkmb,60_ssgjpqz1,61_czhpjzdq,62_hxhpjzdq,63_yqpxqfq1,64_xqtq1m,65_tqzjl1,66_pxqtgdq,67_ssjz1,68_ssjz1,69_rqhjtqdq,70_sslxxxpq,71_dslxpq3c,72_zysllxspqq4c,p_n,e_n,f_n,score) values (:survey_info_id,:child_id,:age,:status,:test_n,:1_zzssldzdqg,:2_zzzyzzzzd,:3_zadswysjdswj,:4_dlzd5s,:5_zlsnwysjdswp,:6_djz5s,:7_djz10szyll,:8_jjz8s,:9_lhpx,:10_pslt,:11_pxlt,:12_tbyd,:13_zwzq,:14_zzzq,:15_ydzq,:16_zzdzmf,:17_wqdszcxqtbfy,:18_wcdszcxptbfy,:19_whdszcxhtbfy,:20_djzyb6syb3s,:21_fysz,:22_dzxz5s,:23_ssbdwjxqz,:24_cxxz,:25_zzx,:26_fwslt,:27_fwxlt,:28_z15cmkdphmsjtxz,:29_lbsyjlt,:30_lbxyjlt,:31_ybsyjlt,:32_ybxyjlt,:33_djz10b,:34_zpdsdtzzx5b,:35_jjtjhgz3m,:36_jhgxz3m,:37_sftt,:38_ydtsjld5cm,:39_xqt10cm,:40_zzlthtjswxt,:41_tsg5cmdtj,:42_tggdjxdlg,:43_zyt3s,:44_xqktyb,:45_lxkmyb10b,:46_lxchb3m,:47_lxxht10c,:48_lxdjt8c,:49_5snz3mjlnlhp,:50_zztqz1,:51_dstwqz1,:52_zstq10c,:53_dsdbbsbx3m,:54_ssdcwtpbx3m,:55_dssgjxqpq1,:56_sssgjxqpq1,:57_ssxxpqz1,:58_dsxxpqz1,:59_dssgjpzkdz3my2mkmb,:60_ssgjpqz1,:61_czhpjzdq,:62_hxhpjzdq,:63_yqpxqfq1,:64_xqtq1m,:65_tqzjl1,:66_pxqtgdq,:67_ssjz1,:68_ssjz1,:69_rqhjtqdq,:70_sslxxxpq,:71_dslxpq3c,:72_zysllxspqq4c,:p_n,:e_n,:f_n,:score)"
	_, err = db.NamedExec(sqlStr1, &sevenSurvey.SurveyCddz)
	fmt.Println("cddz")
	fmt.Println(err)

	sqlStr2 := "Insert into survey_jxdz_info (survey_info_id,child_id,age,status,test_n,1_yzxzwwp,2_ymzszhzzzw,3_ymzhszjswp,4_bwpfrdrqz,5_bxwjfrxpz,6_yhwj,7_tdwjc,8_ls,9_ssrrqzqw,10_pzqxlsw,11_yszadwjkg,12_dkgz,13_yzsnlkjm,14_bxzcrxzb,15_bzwzqdwp,16_zyfs,17_ndwjft,18_yssbwwp,19_jmhj,20_sscwj,21_sspzwj,22_sspzjjjxwj,23_tq,24_nkpg,25_cdb,26_dq2kjm,27_dq7kjm,28_dq1kjm,29_czhdzz,30_cxhdzz,31_cwxdzz,32_cxd,33_cdbzncxzz,34_jxzzfjpz,35_jkz,36_xkz,37_ksdz,38_zz,39_yzxwb,40_ymzszhzzwb,41_zjwbzzsty,42_fhsx,43_fhhx,44_fhyx,45_fhsz,46_fhzfx,47_lx,48_zzdfwnhzx,49_zzdfwnhqx,50_mhqx,51_xnts,52_mxwz,53_ydqkxpn,54_bxpncctz,55_lkhtsbt,56_yjbtpdtx,57_gyzh,58_jz,59_jdzt,60_yzxjz,61_jyx,62_jzfx,63_jfztx,64_yxpcdgznzt,65_bzzfjwjdn,66_yzchx10cm,p_n,e_n,f_n,score) VALUES (:survey_info_id,:child_id,:age,:status,:test_n,:1_yzxzwwp,:2_ymzszhzzzw,:3_ymzhszjswp,:4_bwpfrdrqz,:5_bxwjfrxpz,:6_yhwj,:7_tdwjc,:8_ls,:9_ssrrqzqw,:10_pzqxlsw,:11_yszadwjkg,:12_dkgz,:13_yzsnlkjm,:14_bxzcrxzb,:15_bzwzqdwp,:16_zyfs,:17_ndwjft,:18_yssbwwp,:19_jmhj,:20_sscwj,:21_sspzwj,:22_sspzjjjxwj,:23_tq,:24_nkpg,:25_cdb,:26_dq2kjm,:27_dq7kjm,:28_dq1kjm,:29_czhdzz,:30_cxhdzz,:31_cwxdzz,:32_cxd,:33_cdbzncxzz,:34_jxzzfjpz,:35_jkz,:36_xkz,:37_ksdz,:38_zz,:39_yzxwb,:40_ymzszhzzwb,:41_zjwbzzsty,:42_fhsx,:43_fhhx,:44_fhyx,:45_fhsz,:46_fhzfx,:47_lx,:48_zzdfwnhzx,:49_zzdfwnhqx,:50_mhqx,:51_xnts,:52_mxwz,:53_ydqkxpn,:54_bxpncctz,:55_lkhtsbt,:56_yjbtpdtx,:57_gyzh,:58_jz,:59_jdzt,:60_yzxjz,:61_jyx,:62_jzfx,:63_jfztx,:64_yxpcdgznzt,:65_bzzfjwjdn,:66_yzchx10cm,:p_n,:e_n,:f_n,:score)"
	_, err = db.NamedExec(sqlStr2, &sevenSurvey.SurveyJxdz)
	fmt.Println("gzj")
	fmt.Println(err)

	sqlStr3 := "INSERT INTO survey_yyygt_info (survey_info_id,child_id,age, status, test_n,1_mgjc,2_lbdz,3_sbdz,4_yhdtbystdz,5_ljdtbystdz,6_dsydwzylydzcdfy,7_fbrdsyyqtsy,8_scdyd,9_stdyd,10_ycdyd,11_mfaiu,12_mfoev,13_mfbpm,14_mfhdtg,15_mfkflschnz,16_mfjqxrshzh,17_mfbbmm,18_mfdd,19_mfmgyedfy,20_mfdhdfy,21_mfsl1,22_mfppd,23_mfdwdjsydz,24_mfsxzylw,25_dzjmzyfy,26_ljcjwpmc,27_ljnwt,28_zrstbw,29_zrsw,30_zrjj,31_zrcj,32_zrjtcy,33_zrzj,34_zrdw,35_zrswxwws,36_ljkdddzzl,37_ljfdddzzl,38_ljppsddjzzydbzdz,39_ljlctcbzdz,40_ljbbsjbzdz,41_ljxrhzl,42_ljdx,43_ljds,44_ljcd,45_ljga,46_ljxt,47_ljz,48_ljztybfgx,49_ljssgx,50_ljtjgx,51_ljyggx,52_ljzzgx,53_bdzgzch,54_bdnwt,55_bdcjwp,56_bdstbw,57_bdsg,58_bdb,59_bdgjrl,60_bddx,61_bdds,62_bdz,63_bdcd,64_bdga,65_bdps,66_bdcx,67_bdxtybt,68_bdslddyrlzg,69_bdysddyrhqq,70_bdzjhddzwyjz,71_bdtrhddzwbjzr,72_bdysjxsdjzr,73_zdbdyxhqq,74_bdzdtsywr,75_zdbdyydwj,76_fsyzlgjz,77_fswzdgs,78_mszzfsdsq,79_msyjfsgdsq,p_n,e_n,f_n,score) Values(:survey_info_id, :child_id,:age, :status, :test_n,:1_mgjc,:2_lbdz,:3_sbdz,:4_yhdtbystdz,:5_ljdtbystdz,:6_dsydwzylydzcdfy,:7_fbrdsyyqtsy,:8_scdyd,:9_stdyd,:10_ycdyd,:11_mfaiu,:12_mfoev,:13_mfbpm,:14_mfhdtg,:15_mfkflschnz,:16_mfjqxrshzh,:17_mfbbmm,:18_mfdd,:19_mfmgyedfy,:20_mfdhdfy,:21_mfsl1,:22_mfppd,:23_mfdwdjsydz,:24_mfsxzylw,:25_dzjmzyfy,:26_ljcjwpmc,:27_ljnwt,:28_zrstbw,:29_zrsw,:30_zrjj,:31_zrcj,:32_zrjtcy,:33_zrzj,:34_zrdw,:35_zrswxwws,:36_ljkdddzzl,:37_ljfdddzzl,:38_ljppsddjzzydbzdz,:39_ljlctcbzdz,:40_ljbbsjbzdz,:41_ljxrhzl,:42_ljdx,:43_ljds,:44_ljcd,:45_ljga,:46_ljxt,:47_ljz,:48_ljztybfgx,:49_ljssgx,:50_ljtjgx,:51_ljyggx,:52_ljzzgx,:53_bdzgzch,:54_bdnwt,:55_bdcjwp,:56_bdstbw,:57_bdsg,:58_bdb,:59_bdgjrl,:60_bddx,:61_bdds,:62_bdz,:63_bdcd,:64_bdga,:65_bdps,:66_bdcx,:67_bdxtybt,:68_bdslddyrlzg,:69_bdysddyrhqq,:70_bdzjhddzwyjz,:71_bdtrhddzwbjzr,:72_bdysjxsdjzr,:73_zdbdyxhqq,:74_bdzdtsywr,:75_zdbdyydwj,:76_fsyzlgjz,:77_fswzdgs,:78_mszzfsdsq,:79_msyjfsgdsq,:p_n,:e_n,:f_n,:score)"
	_, err = db.NamedExec(sqlStr3, &sevenSurvey.SurveyYyygt)
	fmt.Println("gzj")
	fmt.Println(err)

	sqlStr4 := "INSERT INTO survey_rzfz_info (survey_info_id,child_id,age, status, test_n,1_azzljcwj,2_zktmlxzczjdstbw,3_zcwjdwdstbw,4_sfsywp,5_brtp,6_zrnhhnh,7_sctpzdxyldbf,8_sctpmc,9_zddzyqdzjhg,10_mbwpjdgx,11_syqz,12_wptpflgn,13_wcjdpt,14_wptpflgn,15_swpd,16_lzzdwppd,17_zcdydehzwdwz,18_ycdplxmg,19_scyzbhnxt,20_fbtpzdsj,21_scsjmc,22_hdjdzmwt,23_xkmj,24_mfjwjfrrq,25_zxycwt,26_sszwsxnwt,27_crqzqcwp,28_zdwpgybffs,29_ayqfzwpsx,30_ayqbfwplw,31_ayqqwpqh,32_pdhlysxt,33_pdhlysby,34_jbysfl,35_scysmc,36_sccjwtys,37_pdszjbxz,38_jbxzfl,39_scxzmc,40_ayqzxhccjxz,41_yyqzcdx,42_zcds,43_zrcd,44_scwjdx,45_qbwtqz,46_brybhzgwt,47_cfd23sz,48_azsnydsz15,49_sc27,50_cs,51_jc26,52_rdsz110,53_cf45,54_jdjaf,55_jdjf,p_n,e_n,f_n,score) Values(:survey_info_id, :child_id,:age, :status, :test_n,:1_azzljcwj,:2_zktmlxzczjdstbw,:3_zcwjdwdstbw,:4_sfsywp,:5_brtp,:6_zrnhhnh,:7_sctpzdxyldbf,:8_sctpmc,:9_zddzyqdzjhg,:10_mbwpjdgx,:11_syqz,:12_wptpflgn,:13_wcjdpt,:14_wptpflgn,:15_swpd,:16_lzzdwppd,:17_zcdydehzwdwz,:18_ycdplxmg,:19_scyzbhnxt,:20_fbtpzdsj,:21_scsjmc,:22_hdjdzmwt,:23_xkmj,:24_mfjwjfrrq,:25_zxycwt,:26_sszwsxnwt,:27_crqzqcwp,:28_zdwpgybffs,:29_ayqfzwpsx,:30_ayqbfwplw,:31_ayqqwpqh,:32_pdhlysxt,:33_pdhlysby,:34_jbysfl,:35_scysmc,:36_sccjwtys,:37_pdszjbxz,:38_jbxzfl,:39_scxzmc,:40_ayqzxhccjxz,:41_yyqzcdx,:42_zcds,:43_zrcd,:44_scwjdx,:45_qbwtqz,:46_brybhzgwt,:47_cfd23sz,:48_azsnydsz15,:49_sc27,:50_cs,:51_jc26,:52_rdsz110,:53_cf45,:54_jdjaf,:55_jdjf,:p_n,:e_n,:f_n,:score)"
	_, err = db.NamedExec(sqlStr4, &sevenSurvey.SurveyRzfz)
	fmt.Println("gzj")
	fmt.Println(err)

	sqlStr5 := "INSERT INTO survey_shzl_info (survey_info_id,child_id,age,status,test_n,1_sxnp,2_tcsw,3_syl,4_hgyl,5_bzhs,6_jjrgt,7_jjygt,8_szkz,9_tcjs,10_czqs,11_swbkz,12_dkrsw,13_ylshdcl,14_kzjsw,15_sswbz,16_rcq,17_rcxy,18_pnpb,19_zbprc,20_rcqlkz,21_rchlkz,22_rchxs,23_fbcsfh,24_dbhqj,25_dbhczqj,26_twwq,27_ttxz,28_tlwz,29_txck,30_twtcs,31_lkll,32_jkdnk,33_ttx,34_cxz,35_cck,36_cwthcs,37_nhdnk,38_ctx,39_chgwz,40_lhll,41_mjcz,42_mjcs,43_xscg,44_mjcl,45_yssy,46_qssk,47_ygsy,48_fzxs,49_ngsmj,50_xmj,51_xl,52_xz,53_stf,54_sjygl,55_ajrs,56_sjaw,57_sjbnc,58_wjgd,59_zdwz,60_gzzdwz,61_wtyj,62_mgs,63_dgs,64_bskm,65_fcj,66_wkssh,67_xw,p_n,e_n,f_n,score) Values(:survey_info_id,:child_id,:age,:status,:test_n,:1_sxnp,:2_tcsw,:3_syl,:4_hgyl,:5_bzhs,:6_jjrgt,:7_jjygt,:8_szkz,:9_tcjs,:10_czqs,:11_swbkz,:12_dkrsw,:13_ylshdcl,:14_kzjsw,:15_sswbz,:16_rcq,:17_rcxy,:18_pnpb,:19_zbprc,:20_rcqlkz,:21_rchlkz,:22_rchxs,:23_fbcsfh,:24_dbhqj,:25_dbhczqj,:26_twwq,:27_ttxz,:28_tlwz,:29_txck,:30_twtcs,:31_lkll,:32_jkdnk,:33_ttx,:34_cxz,:35_cck,:36_cwthcs,:37_nhdnk,:38_ctx,:39_chgwz,:40_lhll,:41_mjcz,:42_mjcs,:43_xscg,:44_mjcl,:45_yssy,:46_qssk,:47_ygsy,:48_fzxs,:49_ngsmj,:50_xmj,:51_xl,:52_xz,:53_stf,:54_sjygl,:55_ajrs,:56_sjaw,:57_sjbnc,:58_wjgd,:59_zdwz,:60_gzzdwz,:61_wtyj,:62_mgs,:63_dgs,:64_bskm,:65_fcj,:66_wkssh,:67_xw,:p_n,:e_n,:f_n,:score)"
	_, err = db.NamedExec(sqlStr5, &sevenSurvey.SurveyShzl)
	fmt.Println("shzl")
	fmt.Println(err)

	sqlStr6 := "INSERT INTO survey_shjw_info (survey_info_id,child_id,age,status,test_n,1_mgzssjdx,2_tp3mt,3_ts3mt,4_msrzjst,5_nmsrzjst,6_az,7_rsjz,8_rsyf,9_zdhdnl,10_zdhdfm,11_pjhh,12_zxbz,13_wxhy,14_wxsyfy,15_wxzksb,16_qqnzjdx,17_dmshjmsr,18_msrjt,19_wcth,20_fx,21_dwhjy,22_wxhywh,23_sslshywh,24_haowh,25_ssayhwh,26_wswh,27_zshwshwh,28_sshayhws,29_nhws,30_sshayhhs,31_bdjstw,32_zdjstw,33_mmlkdmbd,34_mmlkdkmbd,35_mmlkmb,36_mmlkbt,37_mmlkql,38_mmlkbyy,39_mmlkhs,40_mmlkyq,41_mmtmm,42_mmcall,43_mmms,44_nddxgx,45_nddxxxss,46_nhdxsy,47_bscz,p_n,e_n,f_n,score) Values(:survey_info_id,:child_id,:age,:status,:test_n,:1_mgzssjdx,:2_tp3mt,:3_ts3mt,:4_msrzjst,:5_nmsrzjst,:6_az,:7_rsjz,:8_rsyf,:9_zdhdnl,:10_zdhdfm,:11_pjhh,:12_zxbz,:13_wxhy,:14_wxsyfy,:15_wxzksb,:16_qqnzjdx,:17_dmshjmsr,:18_msrjt,:19_wcth,:20_fx,:21_dwhjy,:22_wxhywh,:23_sslshywh,:24_haowh,:25_ssayhwh,:26_wswh,:27_zshwshwh,:28_sshayhws,:29_nhws,:30_sshayhhs,:31_bdjstw,:32_zdjstw,:33_mmlkdmbd,:34_mmlkdkmbd,:35_mmlkmb,:36_mmlkbt,:37_mmlkql,:38_mmlkbyy,:39_mmlkhs,:40_mmlkyq,:41_mmtmm,:42_mmcall,:43_mmms,:44_nddxgx,:45_nddxxxss,:46_nhdxsy,:47_bscz,:p_n,:e_n,:f_n,:score)"
	_, err = db.NamedExec(sqlStr6, &sevenSurvey.SurveyShjw)
	fmt.Println("shjw")
	fmt.Println(err)

	return
}

func DeleteSurveyInfo(cid int, test_n string) (err error) {

	//tx, err := db.Beginx() // 开启事务
	//if err != nil {
	//	fmt.Printf("begin trans failed, err:%v\n", err)
	//	return err
	//}
	//defer func() {
	//	if p := recover(); p != nil {
	//		tx.Rollback()
	//		panic(p) // re-throw panic after Rollback
	//	} else if err != nil {
	//		fmt.Println("rollback")
	//		tx.Rollback() // err is non-nil; don't change it
	//	} else {
	//		err = tx.Commit() // err is nil; if Commit returns error update err
	//		fmt.Println("commit")
	//	}
	//}()

	sqlStr := "Delete from survey_gzj_info where child_id=? and test_n=?"
	_, err = db.Exec(sqlStr, cid, test_n)
	sqlStr1 := "Delete from survey_cddz_info where child_id=? and test_n=?"
	_, err = db.Exec(sqlStr1, cid, test_n)
	sqlStr2 := "Delete from survey_jxdz_info where child_id=? and test_n=?"
	_, err = db.Exec(sqlStr2, cid, test_n)
	sqlStr3 := "Delete from survey_yyygt_info where child_id=? and test_n=?"
	_, err = db.Exec(sqlStr3, cid, test_n)
	sqlStr4 := "Delete from survey_rzfz_info where child_id=? and test_n=?"
	_, err = db.Exec(sqlStr4, cid, test_n)
	sqlStr5 := "Delete from survey_shjw_info where child_id=? and test_n=?"
	_, err = db.Exec(sqlStr5, cid, test_n)
	sqlStr6 := "Delete from survey_shzl_info where child_id=? and test_n=?"
	_, err = db.Exec(sqlStr6, cid, test_n)

	return
}

func SelectBaseAge(survey_id int) (data []int, err error) {
	switch survey_id {
	case 1:
		sqlStr := `select IFNULL(gzj,-1) as gzj from metric_score`
		err = db.Select(&data, sqlStr)
	case 2:
		sqlStr := `select IFNULL(cddz,-1) as cddz from metric_score`
		err = db.Select(&data, sqlStr)
	case 3:
		sqlStr := `select IFNULL(jxdz,-1) as jxdz from metric_score`
		err = db.Select(&data, sqlStr)
	case 4:
		sqlStr := `select IFNULL(yyygt,-1) yyygt from metric_score`
		err = db.Select(&data, sqlStr)
	case 5:
		sqlStr := `select IFNULL(rzfz,-1) rzfz from metric_score`
		err = db.Select(&data, sqlStr)
	case 6:
		sqlStr := `select IFNULL(shjw,-1) shjw from metric_score`
		err = db.Select(&data, sqlStr)
	case 7:
		sqlStr := `select IFNULL(shzl,-1) shzl from metric_score`
		err = db.Select(&data, sqlStr)
	case 8:
		sqlStr := `select IFNULL(fzfs,-1) fzfs from metric_score`
		err = db.Select(&data, sqlStr)
	}

	return
}

//func GetAllInfoGzj(id int, test_n string) (gzj models.SurveyGzj, err error) {
//	sqlStr := `select *
//               from survey_info_id  where child_id=? and test_n=?`
//	err = db.Get(&gzj, sqlStr, id, test_n)
//	return
//}
