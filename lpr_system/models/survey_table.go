package models

// 首页展示孩子表
type ChildInfoTable struct {
	ID          int    `json:"code" db:"child_id" binding:"required"`
	Age         int    `json:"age" db:"age" binding:"required"`
	Name        string `json:"name" db:"name"`
	TestN       string `json:"test_n" db:"test_n" binding:"required"`
	ManagerName string `json:"manager_name" db:"manage_name"`
}

// 七个表符合年龄范围的所有项目
type SelectedSevenTableItems struct {
	ID       int    `json:"code" db:"child_id"`
	Age      int    `json:"age" db:"age" `
	TestN    string `json:"test_n" db:"test_n" `
	Status   int    `json:"status" db:"status" `
	AllItems [7]SelectedTableItems
}

type SelectedTableItems struct {
	TabelName string `json:"survey_name" db:"survey_name" binding:"required"`
	Items     []Item
}

type Item struct {
	SurveySubRange      string `json:"survey_sub_range" db:"survey_sub_range" binding:"required"`
	SurveyProjectCnname string `json:"survey_project_cnname" db:"survey_project_cnname" binding:"required"`
	SurveyProjectEnname string `json:"survey_project_enname" db:"survey_project_enname" binding:"required"`
	Score               string `json:"score" db:"score" binding:"required"`
}

// 首页展示孩子表
type OneSurveyInfo struct {
	Surveyname    string
	AllTimesItems TimesItems
}

type TimesItems struct {
	PN            int    `json:"p_n" db:"p_n" binding:"required"`
	EN            int    `json:"e_n" db:"e_n" binding:"required"`
	FN            int    `json:"f_n" db:"f_n" binding:"required"`
	TestN         string `json:"test_n" db:"test_n" binding:"required"`
	ItemsWithTest []ItemOneSurvey
}

type ItemOneSurvey struct {
	SurveyProjectCnname string `json:"survey_project_cnname" db:"survey_project_cnname" binding:"required"`
	Score               string `json:"score" db:"score" binding:"required"`
}

// 传来表id和孩子id和次数的参数
type SurveyChildID struct {
	SurveyID int    `json:"survey_id"`
	ChildID  int    `json:"child_id" db:"child_id"`
	TestN    string `json:"test_n" db:"test_n" binding:"required"`
}

// 感知觉调查表
type SurveyGzj struct {
	ID      int    `json:"id"     db:"survey_info_id"`
	ChildID int    `json:"child_id" db:"child_id"`
	Age     int    `json:"age" db:"age"`
	Status  int    `json:"status" db:"status"`
	TestN   string `json:"test_n" db:"test_n"`
	P1      string `json:"1_gxcj" db:"1_gxcj"`
	P2      string `json:"2_yscj" db:"2_yscj"`
	P3      string `json:"3_lhzs" db:"3_lhzs"`
	P4      string `json:"4_pdwt" db:"4_pdwt"`
	P5      string `json:"5_ksyd" db:"5_ksyd"`
	P6      string `json:"6_sbmk" db:"6_sbmk"`
	P7      string `json:"7_bryx" db:"7_bryx"`
	P8      string `json:"8_brwp" db:"8_brwp"`
	P9      string `json:"9_brys" db:"9_brys"`
	P10     string `json:"10_brys" db:"10_brys"`
	P11     string `json:"11_btys" db:"11_btys"`
	P12     string `json:"12_txpk" db:"12_txpk"`
	P13     string `json:"13_dwpk" db:"13_dwpk"`
	P14     string `json:"14_wppb" db:"14_wppb"`
	P15     string `json:"15_dwpt" db:"15_dwpt"`
	P16     string `json:"16_jmpd" db:"16_jmpd"`
	P17     string `json:"17_bxxw" db:"17_bxxw"`
	P18     string `json:"18_thxq" db:"18_thxq"`
	P19     string `json:"19_gyy" db:"19_gyy"`
	P20     string `json:"20_jmfy" db:"20_jmfy"`
	P21     string `json:"21_tffy" db:"21_tffy"`
	P22     string `json:"22_tzsy" db:"22_tzsy"`
	P23     string `json:"23_zysy" db:"23_zysy"`
	P24     string `json:"24_ltyc" db:"24_ltyc"`
	P25     string `json:"25_ltmb" db:"25_ltmb"`
	P26     string `json:"26_jmfy " db:"26_jmfy"`
	P27     string `json:"27_ltxb" db:"27_ltxb"`
	P28     string `json:"28_ltsz" db:"28_ltsz"`
	P29     string `json:"29_ltyl" db:"29_ltyl"`
	P30     string `json:"30_fbsy" db:"30_fbsy"`
	P31     string `json:"31_fbsy" db:"31_fbsy"`
	P32     string `json:"32_btyq" db:"32_btyq"`
	P33     string `json:"33_fbls" db:"33_fbls"`
	P34     string `json:"34_pzfy" db:"34_pzfy"`
	P35     string `json:"35_cjfy " db:"35_cjfy"`
	P36     string `json:"36_ysfy" db:"36_ysfy"`
	P37     string `json:"37_cjpd" db:"37_cjpd"`
	P38     string `json:"38_cjtx" db:"38_cjtx"`
	P39     string `json:"39_qfqz" db:"39_qfqz"`
	P40     string `json:"40_qfgs " db:"40_qfgs"`
	P41     string `json:"41_qfxz" db:"41_qfxz"`
	P42     string `json:"42_dxpl" db:"42_dxpl"`
	P43     string `json:"43_qzpl" db:"43_qzpl"`
	P44     string `json:"44_wdfy" db:"44_wdfy"`
	P45     string `json:"45_fbst" db:"45_fbst"`
	P46     string `json:"46_fbxk " db:"46_fbxk"`
	P47     string `json:"47_bblr" db:"47_bblr"`
	P48     string `json:"48_bbnd" db:"48_bbnd"`
	P49     string `json:"49_bbry" db:"49_bbry"`
	P50     string `json:"50_bbhh" db:"50_bbhh"`
	P51     string `json:"51_szqw" db:"51_szqw"`
	P52     string `json:"52_bbxc" db:"52_bbxc"`
	P53     string `json:"53_qwpd" db:"53_qwpd"`
	P54     string `json:"54_scsw" db:"54_scsw"`
	P55     string `json:"55_scsg" db:"55_scsg"`
	PN      int    `json:"p_n" db:"p_n"`
	EN      int    `json:"e_n" db:"e_n"`
	FN      int    `json:"f_n" db:"f_n"`
	Score   int    `json:"score" db:"score"`
}

// 粗大动作调查表
type SurveyCddz struct {
	ID      int    `json:"id"     db:"survey_info_id"`
	ChildID int    `json:"child_id" db:"child_id"`
	Age     int    `json:"age" db:"age"`
	Status  int    `json:"status" db:"status"`
	TestN   string `json:"test_n" db:"test_n"`
	P1      string `json:"1_zzssldzdqg" db:"1_zzssldzdqg"`
	P2      string `json:"2_zzzyzzzzd" db:"2_zzzyzzzzd"`
	P3      string `json:"3_zadswysjdswj" db:"3_zadswysjdswj"`
	P4      string `json:"4_dlzd5s" db:"4_dlzd5s"`
	P5      string `json:"5_zlsnwysjdswp" db:"5_zlsnwysjdswp"`
	P6      string `json:"6_djz5s" db:"6_djz5s"`
	P7      string `json:"7_djz10szyll" db:"7_djz10szyll"`
	P8      string `json:"8_jjz8s" db:"8_jjz8s"`
	P9      string `json:"9_lhpx" db:"9_lhpx"`
	P10     string `json:"10_pslt" db:"10_pslt"`
	P11     string `json:"11_pxlt" db:"11_pxlt"`
	P12     string `json:"12_tbyd" db:"12_tbyd"`
	P13     string `json:"13_zwzq" db:"13_zwzq"`
	P14     string `json:"14_zzzq" db:"14_zzzq"`
	P15     string `json:"15_ydzq" db:"15_ydzq"`
	P16     string `json:"16_zzdzmf" db:"16_zzdzmf"`
	P17     string `json:"17_wqdszcxqtbfy" db:"17_wqdszcxqtbfy"`
	P18     string `json:"18_wcdszcxptbfy" db:"18_wcdszcxptbfy"`
	P19     string `json:"19_whdszcxhtbfy" db:"19_whdszcxhtbfy"`
	P20     string `json:"20_djzyb6syb3s" db:"20_djzyb6syb3s"`
	P21     string `json:"21_fysz" db:"21_fysz"`
	P22     string `json:"22_dzxz5s" db:"22_dzxz5s"`
	P23     string `json:"23_ssbdwjxqz" db:"23_ssbdwjxqz"`
	P24     string `json:"24_cxxz" db:"24_cxxz"`
	P25     string `json:"25_zzx" db:"25_zzx"`
	P26     string `json:"26_fwslt " db:"26_fwslt"`
	P27     string `json:"27_fwxlt" db:"27_fwxlt"`
	P28     string `json:"28_z15cmkdphmsjtxz" db:"28_z15cmkdphmsjtxz"`
	P29     string `json:"29_lbsyjlt(bfw)" db:"29_lbsyjlt"`
	P30     string `json:"30_lbxyjlt(bfw)" db:"30_lbxyjlt"`
	P31     string `json:"31_ybsyjlt" db:"31_ybsyjlt"`
	P32     string `json:"32_ybxyjlt" db:"32_ybxyjlt"`
	P33     string `json:"33_djz10b" db:"33_djz10b"`
	P34     string `json:"34_zpdsdtzzx5b" db:"34_zpdsdtzzx5b"`
	P35     string `json:"35_jjtjhgz3m" db:"35_jjtjhgz3m"`
	P36     string `json:"36_jhgxz3m" db:"36_jhgxz3m"`
	P37     string `json:"37_sftt" db:"37_sftt"`
	P38     string `json:"38_ydtsjld5cm" db:"38_ydtsjld5cm"`
	P39     string `json:"39_xqt10cm" db:"39_xqt10cm"`
	P40     string `json:"40_zzlthtjswxt" db:"40_zzlthtjswxt"`
	P41     string `json:"41_tsg5cmdtj" db:"41_tsg5cmdtj"`
	P42     string `json:"42_tggdjxdlg" db:"42_tggdjxdlg"`
	P43     string `json:"43_zyt3s" db:"43_zyt3s"`
	P44     string `json:"44_xqktyb" db:"44_xqktyb"`
	P45     string `json:"45_lxkmyb10b" db:"45_lxkmyb10b"`
	P46     string `json:"46_lxchb3m" db:"46_lxchb3m"`
	P47     string `json:"47_lxxht10c" db:"47_lxxht10c"`
	P48     string `json:"48_lxdjt8c" db:"48_lxdjt8c"`
	P49     string `json:"49_5snz3mjlnlhp" db:"49_5snz3mjlnlhp"`
	P50     string `json:"50_zztqz1.5mcy" db:"50_zztqz1"`
	P51     string `json:"51_dstwqz1.8myk40cmdqmn" db:"51_dstwqz1"`
	P52     string `json:"52_zstq10c" db:"52_zstq10c"`
	P53     string `json:"53_dsdbbsbx3m" db:"53_dsdbbsbx3m"`
	P54     string `json:"54_ssdcwtpbx3m" db:"54_ssdcwtpbx3m"`
	P55     string `json:"55_dssgjxqpq1" db:"55_dssgjxqpq1"`
	P56     string `json:"56_sssgjxqpq1" db:"56_sssgjxqpq1"`
	P57     string `json:"57_ssxxpqz1" db:"57_ssxxpqz1"`
	P58     string `json:"58_dsxxpqz1" db:"58_dsxxpqz1"`
	P59     string `json:"59_dssgjpzkdz3my2mkmb" db:"59_dssgjpzkdz3my2mkmb"`
	P60     string `json:"60_ssgjpqz1" db:"60_ssgjpqz1"`
	P61     string `json:"61_czhpjzdq" db:"61_czhpjzdq"`
	P62     string `json:"62_hxhpjzdq" db:"62_hxhpjzdq"`
	P63     string `json:"63_yqpxqfq1.5m" db:"63_yqpxqfq1"`
	P64     string `json:"64_xqtq1m" db:"64_xqtq1m"`
	P65     string `json:"65_tqzjl1.8myk70cmdmb" db:"65_tqzjl1"`
	P66     string `json:"66_pxqtgdq" db:"66_pxqtgdq"`
	P67     string `json:"67_ssjz1.5mycpldq" db:"67_ssjz1"`
	P68     string `json:"68_ssjz1.5mycdsthldq" db:"68_ssjz1"`
	P69     string `json:"69_rqhjtqdq" db:"69_rqhjtqdq"`
	P70     string `json:"70_sslxxxpq" db:"70_sslxxxpq"`
	P71     string `json:"71_dslxpq3c" db:"71_dslxpq3c"`
	P72     string `json:"72_zysllxspqq4c" db:"72_zysllxspqq4c"`
	PN      int    `json:"p_n" db:"p_n"`
	EN      int    `json:"e_n" db:"e_n"`
	FN      int    `json:"f_n" db:"f_n"`
	Score   int    `json:"score" db:"score"`
}

// 粗大动作调查表
type SurveyJxdz struct {
	ID      int    `json:"id"     db:"survey_info_id"`
	ChildID int    `json:"child_id" db:"child_id"`
	Age     int    `json:"age" db:"age"`
	Status  int    `json:"status" db:"status"`
	TestN   string `json:"test_n" db:"test_n"`
	P1      string `json:"1_yzxzwwp" db:"1_yzxzwwp"`
	P2      string `json:"2_ymzszhzzzw" db:"2_ymzszhzzzw"`
	P3      string `json:"3_ymzhszjswp" db:"3_ymzhszjswp"`
	P4      string `json:"4_bwpfrdrqz" db:"4_bwpfrdrqz"`
	P5      string `json:"5_bxwjfrxpz" db:"5_bxwjfrxpz"`
	P6      string `json:"6_yhwj" db:"6_yhwj"`
	P7      string `json:"7_tdwjc" db:"7_tdwjc"`
	P8      string `json:"8_ls" db:"8_ls"`
	P9      string `json:"9_ssrrqzqw" db:"9_ssrrqzqw"`
	P10     string `json:"10_pzqxlsw" db:"10_pzqxlsw"`
	P11     string `json:"11_yszadwjkg" db:"11_yszadwjkg"`
	P12     string `json:"12_dkgz" db:"12_dkgz"`
	P13     string `json:"13_yzsnlkjm" db:"13_yzsnlkjm"`
	P14     string `json:"14_bxzcrxzb" db:"14_bxzcrxzb"`
	P15     string `json:"15_bzwzqdwp" db:"15_bzwzqdwp"`
	P16     string `json:"16_zyfs" db:"16_zyfs"`
	P17     string `json:"17_ndwjft" db:"17_ndwjft"`
	P18     string `json:"18_yssbwwp" db:"18_yssbwwp"`
	P19     string `json:"19_jmhj" db:"19_jmhj"`
	P20     string `json:"20_sscwj" db:"20_sscwj"`
	P21     string `json:"21_sspzwj" db:"21_sspzwj"`
	P22     string `json:"22_sspzjjjxwj" db:"22_sspzjjjxwj"`
	P23     string `json:"23_tq" db:"23_tq"`
	P24     string `json:"24_nkpg" db:"24_nkpg"`
	P25     string `json:"25_cdb" db:"25_cdb"`
	P26     string `json:"26_dq2kjm" db:"26_dq2kjm"`
	P27     string `json:"27_dq7kjm" db:"27_dq7kjm"`
	P28     string `json:"28_dq1kjm" db:"28_dq1kjm"`
	P29     string `json:"29_czhdzz" db:"29_czhdzz"`
	P30     string `json:"30_cxhdzz" db:"30_cxhdzz"`
	P31     string `json:"31_cwxdzz" db:"31_cwxdzz"`
	P32     string `json:"32_cxd" db:"32_cxd"`
	P33     string `json:"33_cdbzncxzz" db:"33_cdbzncxzz"`
	P34     string `json:"34_jxzzfjpz" db:"34_jxzzfjpz"`
	P35     string `json:"35_jkz" db:"35_jkz"`
	P36     string `json:"36_xkz" db:"36_xkz"`
	P37     string `json:"37_ksdz" db:"37_ksdz"`
	P38     string `json:"38_zz" db:"38_zz"`
	P39     string `json:"39_yzxwb" db:"39_yzxwb"`
	P40     string `json:"40_ymzszhzzwb" db:"40_ymzszhzzwb"`
	P41     string `json:"41_zjwbzzsty" db:"41_zjwbzzsty"`
	P42     string `json:"42_fhsx" db:"42_fhsx"`
	P43     string `json:"43_fhhx" db:"43_fhhx"`
	P44     string `json:"44_fhyx" db:"44_fhyx"`
	P45     string `json:"45_fhsz" db:"45_fhsz"`
	P46     string `json:"46_fhzfx" db:"46_fhzfx"`
	P47     string `json:"47_lx" db:"47_lx"`
	P48     string `json:"48_zzdfwnhzx" db:"48_zzdfwnhzx"`
	P49     string `json:"49_zzdfwnhqx" db:"49_zzdfwnhqx"`
	P50     string `json:"50_mhqx" db:"50_mhqx"`
	P51     string `json:"51_xnts" db:"51_xnts"`
	P52     string `json:"52_mxwz" db:"52_mxwz"`
	P53     string `json:"53_ydqkxpn" db:"53_ydqkxpn"`
	P54     string `json:"54_bxpncctz" db:"54_bxpncctz"`
	P55     string `json:"55_lkhtsbt" db:"55_lkhtsbt"`
	P56     string `json:"56_yjbtpdtx" db:"56_yjbtpdtx"`
	P57     string `json:"57_gyzh" db:"57_gyzh"`
	P58     string `json:"58_jz" db:"58_jz"`
	P59     string `json:"59_jdzt" db:"59_jdzt"`
	P60     string `json:"60_yzxjz" db:"60_yzxjz"`
	P61     string `json:"61_jyx" db:"61_jyx"`
	P62     string `json:"62_jzfx" db:"62_jzfx"`
	P63     string `json:"63_jfztx" db:"63_jfztx"`
	P64     string `json:"64_yxpcdgznzt" db:"64_yxpcdgznzt"`
	P65     string `json:"65_bzzfjwjdn" db:"65_bzzfjwjdn"`
	P66     string `json:"66_yzchx10cm" db:"66_yzchx10cm"`
	PN      int    `json:"p_n" db:"p_n"`
	EN      int    `json:"e_n" db:"e_n"`
	FN      int    `json:"f_n" db:"f_n"`
	Score   int    `json:"score" db:"score"`
}

// 语言与沟通调查表
type SurveyYyygt struct {
	ID      int    `json:"id"     db:"survey_info_id"`
	ChildID int    `json:"child_id" db:"child_id"`
	Age     int    `json:"age" db:"age"`
	Status  int    `json:"status" db:"status"`
	TestN   string `json:"test_n" db:"test_n"`
	P1      string `json:"1_mgjc" db:"1_mgjc"`
	P2      string `json:"2_lbdz" db:"2_lbdz"`
	P3      string `json:"3_sbdz" db:"3_sbdz"`
	P4      string `json:"4_yhdtbystdz" db:"4_yhdtbystdz"`
	P5      string `json:"5_ljdtbystdz" db:"5_ljdtbystdz"`
	P6      string `json:"6_dsydwzylydzcdfy" db:"6_dsydwzylydzcdfy"`
	P7      string `json:"7_fbrdsyyqtsy" db:"7_fbrdsyyqtsy"`
	P8      string `json:"8_scdyd" db:"8_scdyd"`
	P9      string `json:"9_stdyd" db:"9_stdyd"`
	P10     string `json:"10_ycdyd" db:"10_ycdyd"`
	P11     string `json:"11_mfaiu" db:"11_mfaiu"`
	P12     string `json:"12_mfoev" db:"12_mfoev"`
	P13     string `json:"13_mfbpm" db:"13_mfbpm"`
	P14     string `json:"14_mfhdtg" db:"14_mfhdtg"`
	P15     string `json:"15_mfkflschnz" db:"15_mfkflschnz"`
	P16     string `json:"16_mfjqxrshzh" db:"16_mfjqxrshzh"`
	P17     string `json:"17_mfbbmm" db:"17_mfbbmm"`
	P18     string `json:"18_mfdd" db:"18_mfdd"`
	P19     string `json:"19_mfmgyedfy" db:"19_mfmgyedfy"`
	P20     string `json:"20_mfdhdfy" db:"20_mfdhdfy"`
	P21     string `json:"21_mfsl1-10dfy" db:"21_mfsl1"`
	P22     string `json:"22_mfppd" db:"22_mfppd"`
	P23     string `json:"23_mfdwdjsydz" db:"23_mfdwdjsydz"`
	P24     string `json:"24_mfsxzylw" db:"24_mfsxzylw"`
	P25     string `json:"25_dzjmzyfy" db:"25_dzjmzyfy"`
	P26     string `json:"26_ljcjwpmc " db:"26_ljcjwpmc"`
	P27     string `json:"27_ljnwt" db:"27_ljnwt"`
	P28     string `json:"28_zrstbw" db:"28_zrstbw"`
	P29     string `json:"29_zrsw" db:"29_zrsw"`
	P30     string `json:"30_zrjj" db:"30_zrjj"`
	P31     string `json:"31_zrcj" db:"31_zrcj"`
	P32     string `json:"32_zrjtcy" db:"32_zrjtcy"`
	P33     string `json:"33_zrzj" db:"33_zrzj"`
	P34     string `json:"34_zrdw" db:"34_zrdw"`
	P35     string `json:"35_zrswxwws " db:"35_zrswxwws"`
	P36     string `json:"36_ljkdddzzl" db:"36_ljkdddzzl"`
	P37     string `json:"37_ljfdddzzl" db:"37_ljfdddzzl"`
	P38     string `json:"38_ljppsddjzzydbzdz" db:"38_ljppsddjzzydbzdz"`
	P39     string `json:"39_ljlctcbzdz" db:"39_ljlctcbzdz"`
	P40     string `json:"40_ljbbsjbzdz " db:"40_ljbbsjbzdz"`
	P41     string `json:"41_ljxrhzl" db:"41_ljxrhzl"`
	P42     string `json:"42_ljdx" db:"42_ljdx"`
	P43     string `json:"43_ljds" db:"43_ljds"`
	P44     string `json:"44_ljcd" db:"44_ljcd"`
	P45     string `json:"45_ljga" db:"45_ljga"`
	P46     string `json:"46_ljxt " db:"46_ljxt"`
	P47     string `json:"47_ljz" db:"47_ljz"`
	P48     string `json:"48_ljztybfgx" db:"48_ljztybfgx"`
	P49     string `json:"49_ljssgx" db:"49_ljssgx"`
	P50     string `json:"50_ljtjgx" db:"50_ljtjgx"`
	P51     string `json:"51_ljyggx" db:"51_ljyggx"`
	P52     string `json:"52_ljzzgx" db:"52_ljzzgx"`
	P53     string `json:"53_bdzgzch" db:"53_bdzgzch"`
	P54     string `json:"54_bdnwt" db:"54_bdnwt"`
	P55     string `json:"55_bdcjwp" db:"55_bdcjwp"`
	P56     string `json:"56_bdstbw" db:"56_bdstbw"`
	P57     string `json:"57_bdsg" db:"57_bdsg"`
	P58     string `json:"58_bdb" db:"58_bdb"`
	P59     string `json:"59_bdgjrl" db:"59_bdgjrl"`
	P60     string `json:"60_bddx" db:"60_bddx"`
	P61     string `json:"61_bdds" db:"61_bdds"`
	P62     string `json:"62_bdz" db:"62_bdz"`
	P63     string `json:"63_bdcd" db:"63_bdcd"`
	P64     string `json:"64_bdga" db:"64_bdga"`
	P65     string `json:"65_bdps" db:"65_bdps"`
	P66     string `json:"66_bdcx" db:"66_bdcx"`
	P67     string `json:"67_bdxtybt" db:"67_bdxtybt"`
	P68     string `json:"68_bdslddyrlzg" db:"68_bdslddyrlzg"`
	P69     string `json:"69_bdysddyrhqq" db:"69_bdysddyrhqq"`
	P70     string `json:"70_bdzjhddzwyjz" db:"70_bdzjhddzwyjz"`
	P71     string `json:"71_bdtrhddzwbjzr" db:"71_bdtrhddzwbjzr"`
	P72     string `json:"72_bdysjxsdjzr" db:"72_bdysjxsdjzr"`
	P73     string `json:"73_zdbdyxhqq" db:"73_zdbdyxhqq"`
	P74     string `json:"74_bdzdtsywr" db:"74_bdzdtsywr"`
	P75     string `json:"75_zdbdyydwj" db:"75_zdbdyydwj"`
	P76     string `json:"76_fsyzlgjz" db:"76_fsyzlgjz"`
	P77     string `json:"77_fswzdgs" db:"77_fswzdgs"`
	P78     string `json:"78_mszzfsdsq" db:"78_mszzfsdsq"`
	P79     string `json:"79_msyjfsgdsq" db:"79_msyjfsgdsq"`
	PN      int    `json:"p_n" db:"p_n"`
	EN      int    `json:"e_n" db:"e_n"`
	FN      int    `json:"f_n" db:"f_n"`
	Score   int    `json:"score" db:"score"`
}

// 认知发展调查表
type SurveyRzfz struct {
	ID      int    `json:"id"     db:"survey_info_id"`
	ChildID int    `json:"child_id" db:"child_id"`
	Age     int    `json:"age" db:"age"`
	Status  int    `json:"status" db:"status"`
	TestN   string `json:"test_n" db:"test_n"`
	P1      string `json:"1_azzljcwj" db:"1_azzljcwj"`
	P2      string `json:"2_zktmlxzczjdstbw" db:"2_zktmlxzczjdstbw"`
	P3      string `json:"3_zcwjdwdstbw" db:"3_zcwjdwdstbw"`
	P4      string `json:"4_sfsywp" db:"4_sfsywp"`
	P5      string `json:"5_brtp" db:"5_brtp"`
	P6      string `json:"6_zrnhhnh" db:"6_zrnhhnh"`
	P7      string `json:"7_sctpzdxyldbf" db:"7_sctpzdxyldbf"`
	P8      string `json:"8_sctpmc" db:"8_sctpmc"`
	P9      string `json:"9_zddzyqdzjhg" db:"9_zddzyqdzjhg"`
	P10     string `json:"10_mbwpjdgx" db:"10_mbwpjdgx"`
	P11     string `json:"11_syqz" db:"11_syqz"`
	P12     string `json:"12_wptpflgn" db:"12_wptpflgn"`
	P13     string `json:"13_wcjdpt" db:"13_wcjdpt"`
	P14     string `json:"14_wptpflgn" db:"14_wptpflgn"`
	P15     string `json:"15_swpd" db:"15_swpd"`
	P16     string `json:"16_lzzdwppd" db:"16_lzzdwppd"`
	P17     string `json:"17_zcdydehzwdwz" db:"17_zcdydehzwdwz"`
	P18     string `json:"18_ycdplxmg" db:"18_ycdplxmg"`
	P19     string `json:"19_scyzbhnxt" db:"19_scyzbhnxt"`
	P20     string `json:"20_fbtpzdsj" db:"20_fbtpzdsj"`
	P21     string `json:"21_scsjmc" db:"21_scsjmc"`
	P22     string `json:"22_hdjdzmwt" db:"22_hdjdzmwt"`
	P23     string `json:"23_xkmj" db:"23_xkmj"`
	P24     string `json:"24_mfjwjfrrq" db:"24_mfjwjfrrq"`
	P25     string `json:"25_zxycwt" db:"25_zxycwt"`
	P26     string `json:"26_sszwsxnwt " db:"26_sszwsxnwt"`
	P27     string `json:"27_crqzqcwp" db:"27_crqzqcwp"`
	P28     string `json:"28_zdwpgybffs" db:"28_zdwpgybffs"`
	P29     string `json:"29_ayqfzwpsx" db:"29_ayqfzwpsx"`
	P30     string `json:"30_ayqbfwplw" db:"30_ayqbfwplw"`
	P31     string `json:"31_ayqqwpqh" db:"31_ayqqwpqh"`
	P32     string `json:"32_pdhlysxt" db:"32_pdhlysxt"`
	P33     string `json:"33_pdhlysby" db:"33_pdhlysby"`
	P34     string `json:"34_jbysfl" db:"34_jbysfl"`
	P35     string `json:"35_scysmc " db:"35_scysmc"`
	P36     string `json:"36_sccjwtys" db:"36_sccjwtys"`
	P37     string `json:"37_pdszjbxz" db:"37_pdszjbxz"`
	P38     string `json:"38_jbxzfl" db:"38_jbxzfl"`
	P39     string `json:"39_scxzmc" db:"39_scxzmc"`
	P40     string `json:"40_ayqzxhccjxz " db:"40_ayqzxhccjxz"`
	P41     string `json:"41_yyqzcdx" db:"41_yyqzcdx"`
	P42     string `json:"42_zcds" db:"42_zcds"`
	P43     string `json:"43_zrcd" db:"43_zrcd"`
	P44     string `json:"44_scwjdx" db:"44_scwjdx"`
	P45     string `json:"45_qbwtqz" db:"45_qbwtqz"`
	P46     string `json:"46_brybhzgwt " db:"46_brybhzgwt"`
	P47     string `json:"47_cfd23sz" db:"47_cfd23sz"`
	P48     string `json:"48_azsnydsz15" db:"48_azsnydsz15"`
	P49     string `json:"49_sc27" db:"49_sc27"`
	P50     string `json:"50_cs" db:"50_cs"`
	P51     string `json:"51_jc26" db:"51_jc26"`
	P52     string `json:"52_rdsz110" db:"52_rdsz110"`
	P53     string `json:"53_cf45" db:"53_cf45"`
	P54     string `json:"54_jdjaf" db:"54_jdjaf"`
	P55     string `json:"55_jdjf" db:"55_jdjf"`
	PN      int    `json:"p_n" db:"p_n"`
	EN      int    `json:"e_n" db:"e_n"`
	FN      int    `json:"f_n" db:"f_n"`
	Score   int    `json:"score" db:"score"`
}

// 社会交往调查表
type SurveyShjw struct {
	ID      int    `json:"id"     db:"survey_info_id"`
	ChildID int    `json:"child_id" db:"child_id"`
	Age     int    `json:"age" db:"age"`
	Status  int    `json:"status" db:"status"`
	TestN   string `json:"test_n" db:"test_n"`
	P1      string `json:"1_mgzssjdx" db:"1_mgzssjdx"`
	P2      string `json:"2_tp3mt" db:"2_tp3mt"`
	P3      string `json:"3_ts3mt" db:"3_ts3mt"`
	P4      string `json:"4_msrzjst" db:"4_msrzjst"`
	P5      string `json:"5_nmsrzjst" db:"5_nmsrzjst"`
	P6      string `json:"6_az" db:"6_az"`
	P7      string `json:"7_rsjz" db:"7_rsjz"`
	P8      string `json:"8_rsyf" db:"8_rsyf"`
	P9      string `json:"9_zdhdnl" db:"9_zdhdnl"`
	P10     string `json:"10_zdhdfm" db:"10_zdhdfm"`
	P11     string `json:"11_pjhh" db:"11_pjhh"`
	P12     string `json:"12_zxbz" db:"12_zxbz"`
	P13     string `json:"13_wxhy" db:"13_wxhy"`
	P14     string `json:"14_wxsyfy" db:"14_wxsyfy"`
	P15     string `json:"15_wxzksb" db:"15_wxzksb"`
	P16     string `json:"16_qqnzjdx" db:"16_qqnzjdx"`
	P17     string `json:"17_dmshjmsr" db:"17_dmshjmsr"`
	P18     string `json:"18_msrjt" db:"18_msrjt"`
	P19     string `json:"19_wcth" db:"19_wcth"`
	P20     string `json:"20_fx" db:"20_fx"`
	P21     string `json:"21_dwhjy" db:"21_dwhjy"`
	P22     string `json:"22_wxhywh" db:"22_wxhywh"`
	P23     string `json:"23_sslshywh" db:"23_sslshywh"`
	P24     string `json:"24_haowh" db:"24_haowh"`
	P25     string `json:"25_ssayhwh" db:"25_ssayhwh"`
	P26     string `json:"26_wswh " db:"26_wswh"`
	P27     string `json:"27_zshwshwh" db:"27_zshwshwh"`
	P28     string `json:"28_sshayhws" db:"28_sshayhws"`
	P29     string `json:"29_nhws" db:"29_nhws"`
	P30     string `json:"30_sshayhhs" db:"30_sshayhhs"`
	P31     string `json:"31_bdjstw" db:"31_bdjstw"`
	P32     string `json:"32_zdjstw" db:"32_zdjstw"`
	P33     string `json:"33_mmlkdmbd" db:"33_mmlkdmbd"`
	P34     string `json:"34_mmlkdkmbd" db:"34_mmlkdkmbd"`
	P35     string `json:"35_mmlkmb " db:"35_mmlkmb"`
	P36     string `json:"36_mmlkbt" db:"36_mmlkbt"`
	P37     string `json:"37_mmlkql" db:"37_mmlkql"`
	P38     string `json:"38_mmlkbyy" db:"38_mmlkbyy"`
	P39     string `json:"39_mmlkhs" db:"39_mmlkhs"`
	P40     string `json:"40_mmlkyq " db:"40_mmlkyq"`
	P41     string `json:"41_mmtmm" db:"41_mmtmm"`
	P42     string `json:"42_mmcall" db:"42_mmcall"`
	P43     string `json:"43_mmms" db:"43_mmms"`
	P44     string `json:"44_nddxgx" db:"44_nddxgx"`
	P45     string `json:"45_nddxxxss" db:"45_nddxxxss"`
	P46     string `json:"46_nhdxsy " db:"46_nhdxsy"`
	P47     string `json:"47_bscz " db:"47_bscz"`
	PN      int    `json:"p_n" db:"p_n"`
	EN      int    `json:"e_n" db:"e_n"`
	FN      int    `json:"f_n" db:"f_n"`
	Score   int    `json:"score" db:"score"`
}

// 生活自理调查表
type SurveyShzl struct {
	ID      int    `json:"id"     db:"survey_info_id"`
	ChildID int    `json:"child_id" db:"child_id"`
	Age     int    `json:"age" db:"age"`
	Status  int    `json:"status" db:"status"`
	TestN   string `json:"test_n" db:"test_n"`
	P1      string `json:"1_sxnp" db:"1_sxnp"`
	P2      string `json:"2_tcsw" db:"2_tcsw"`
	P3      string `json:"3_syl" db:"3_syl"`
	P4      string `json:"4_hgyl" db:"4_hgyl"`
	P5      string `json:"5_bzhs" db:"5_bzhs"`
	P6      string `json:"6_jjrgt" db:"6_jjrgt"`
	P7      string `json:"7_jjygt" db:"7_jjygt"`
	P8      string `json:"8_szkz" db:"8_szkz"`
	P9      string `json:"9_tcjs" db:"9_tcjs"`
	P10     string `json:"10_czqs" db:"10_czqs"`
	P11     string `json:"11_swbkz" db:"11_swbkz"`
	P12     string `json:"12_dkrsw" db:"12_dkrsw"`
	P13     string `json:"13_ylshdcl" db:"13_ylshdcl"`
	P14     string `json:"14_kzjsw" db:"14_kzjsw"`
	P15     string `json:"15_sswbz" db:"15_sswbz"`
	P16     string `json:"16_rcq" db:"16_rcq"`
	P17     string `json:"17_rcxy" db:"17_rcxy"`
	P18     string `json:"18_pnpb" db:"18_pnpb"`
	P19     string `json:"19_zbprc" db:"19_zbprc"`
	P20     string `json:"20_rcqlkz" db:"20_rcqlkz"`
	P21     string `json:"21_rchlkz" db:"21_rchlkz"`
	P22     string `json:"22_rchxs" db:"22_rchxs"`
	P23     string `json:"23_fbcsfh" db:"23_fbcsfh"`
	P24     string `json:"24_dbhqj" db:"24_dbhqj"`
	P25     string `json:"25_dbhczqj" db:"25_dbhczqj"`
	P26     string `json:"26_twwq " db:"26_twwq"`
	P27     string `json:"27_ttxz" db:"27_ttxz"`
	P28     string `json:"28_tlwz" db:"28_tlwz"`
	P29     string `json:"29_txck" db:"29_txck"`
	P30     string `json:"30_twtcs" db:"30_twtcs"`
	P31     string `json:"31_lkll" db:"31_lkll"`
	P32     string `json:"32_jkdnk" db:"32_jkdnk"`
	P33     string `json:"33_ttx" db:"33_ttx"`
	P34     string `json:"34_cxz" db:"34_cxz"`
	P35     string `json:"35_cck " db:"35_cck"`
	P36     string `json:"36_cwthcs" db:"36_cwthcs"`
	P37     string `json:"37_nhdnk" db:"37_nhdnk"`
	P38     string `json:"38_ctx" db:"38_ctx"`
	P39     string `json:"39_chgwz" db:"39_chgwz"`
	P40     string `json:"40_lhll " db:"40_lhll"`
	P41     string `json:"41_mjcz" db:"41_mjcz"`
	P42     string `json:"42_mjcs" db:"42_mjcs"`
	P43     string `json:"43_xscg" db:"43_xscg"`
	P44     string `json:"44_mjcl" db:"44_mjcl"`
	P45     string `json:"45_yssy" db:"45_yssy"`
	P46     string `json:"46_qssk " db:"46_qssk"`
	P47     string `json:"47_ygsy" db:"47_ygsy"`
	P48     string `json:"48_fzxs" db:"48_fzxs"`
	P49     string `json:"49_ngsmj" db:"49_ngsmj"`
	P50     string `json:"50_xmj" db:"50_xmj"`
	P51     string `json:"51_xl" db:"51_xl"`
	P52     string `json:"52_xz" db:"52_xz"`
	P53     string `json:"53_stf" db:"53_stf"`
	P54     string `json:"54_sjygl" db:"54_sjygl"`
	P55     string `json:"55_ajrs" db:"55_ajrs"`
	P56     string `json:"56_sjaw" db:"56_sjaw"`
	P57     string `json:"57_sjbnc" db:"57_sjbnc"`
	P58     string `json:"58_wjgd" db:"58_wjgd"`
	P59     string `json:"59_zdwz" db:"59_zdwz"`
	P60     string `json:"60_gzzdwz" db:"60_gzzdwz"`
	P61     string `json:"61_wtyj" db:"61_wtyj"`
	P62     string `json:"62_mgs" db:"62_mgs"`
	P63     string `json:"63_dgs" db:"63_dgs"`
	P64     string `json:"64_bskm" db:"64_bskm"`
	P65     string `json:"65_fcj" db:"65_fcj"`
	P66     string `json:"66_wkssh" db:"66_wkssh"`
	P67     string `json:"67_xw" db:"67_xw"`
	PN      int    `json:"p_n" db:"p_n"`
	EN      int    `json:"e_n" db:"e_n"`
	FN      int    `json:"f_n" db:"f_n"`
	Score   int    `json:"score" db:"score"`
}

type SevenSurvey struct {
	SurveyGzj
	SurveyCddz
	SurveyJxdz
	SurveyYyygt
	SurveyRzfz
	SurveyShjw
	SurveyShzl
}
