package qitApi

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

func TestRegexp(t *testing.T) {
	// 示例数据，假设这是一个字符串
	data := `










<!DOCTYPE html>
<html>
<head>
	<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1"/>
	<meta charset="utf-8"/>

	<title>学籍信息</title>
	<link rel="shortcut icon" href="/img/icon/favicon.ico" type="image/x-icon"/>
	<meta name="description" content="overview &amp; stats"/>
	
	<!-- bootstrap & fontawesome -->
	<!--     <link rel="stylesheet" type="text/css" href="/static/template.css"/> -->

	<link rel="stylesheet" href="/assets/css/bootstrap.min.css"/>




	<link rel="stylesheet" href="/assets/font-awesome/4.1.0/css/font-awesome.min.css"/>

	<!-- page specific plugin styles -->
	<link rel="stylesheet" href="/assets/css/jquery-ui.min.css" type="text/css"/>
	
	<link rel="stylesheet" href="/assets/css/chosen.css"/>
	
	
	
	<!-- page specific plugin styles -->

	<!-- jquery ui grid -->
	
	<!-- jquery ui grid -->

	<!-- text fonts -->
	
	<link rel="stylesheet" href="/assets/css/fonts-woff2.css"/>

	<!-- ace styles -->
	<link rel="stylesheet" href="/assets/css/ace.min.css" id="main-ace-style"/>

	<!--[if lte IE 9]>
	<link rel="stylesheet" href="/assets/css/ace-part2.min.css"/>
	<![endif]-->
	
	

	<!--[if lte IE 9]>
	<link rel="stylesheet" href="/assets/css/ace-ie.min.css"/>
	<![endif]-->

	<!-- inline styles related to this page -->

	<!-- ace settings handler -->
	<script src="/assets/js/ace-extra.min.js"></script>

	<!-- HTML5shiv and Respond.js for IE8 to support HTML5 elements and media queries -->

	<!--[if lte IE 8]>
	<script src="/assets/js/html5shiv.min.js"></script>
	<script src="/assets/js/respond.min.js"></script>
	<![endif]-->
	<!-- basic scripts -->

	<!--[if !IE]> -->
	<script src="/js/jQuery/jquery-3.4.1.min.js"></script>
	<script src="/js/jQuery/jquery-migrate-3.1.0.min.js"></script>
	<!-- <![endif]-->

	<!--[if IE]>
	<script src="/js/jQuery/jquery-3.4.1.min.js"></script>
	<script src="/js/jQuery/jquery-migrate-1.4.1.min.js"></script>
	<![endif]-->

	<script type="text/javascript">
		if ('ontouchstart' in document.documentElement) document.write("<script src='/assets/js/jquery.mobile.custom.min.js'>" + "<" + "/script>");
	</script>


	<!-- page specific plugin scripts -->
	<script type="text/javascript">
		$.ajaxPrefilter(function (s) {
			if (s.crossDomain) {
				s.contents.javascript = false;
			}
		});
	</script>

	<!--[if lte IE 8]>
	<script src="/assets/js/excanvas.min.js"></script>
	<![endif]-->

	
	<script type="text/javascript" src="/js/My97DatePicker/WdatePicker.js"></script>
	
	<script type="text/javascript" src="/assets/js/chosen.jquery.min.js"></script>

	<!-- ace scripts -->
	<script src="/assets/js/ace-elements.min.js"></script>
	<script src="/assets/js/ace.min.js"></script>

	

	<script type="text/javascript" src="/js/jQuery/jquery.cookie.js"></script>

	<!-- jqgrid -->
	
	<!-- jqgrid -->

	<!-- commoncss -->
	<link rel="stylesheet" href="/css/commons/commoncss.css" type="text/css"/>
	<!-- commoncss -->

	<!-- alert -->
	<script type="text/javascript" src="/assets/js/jquery-ui.min.js"></script>
	
	<script type="text/javascript" src="/js/customjs/pagination.js"></script>
	<script type="text/javascript" src="/js/customjs/urpalert.js"></script>
	<script type="text/javascript" src="/assets/layer/layer.js"></script>

	<script src="/assets/bootstrap/3.2.0/js/bootstrap.min.js"></script>
	<!-- alert -->

	<!-- light bootstrap 两个图表-->
	
	<!-- light bootstrap 两个图表-->
	
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
	
	<link rel="stylesheet" href="/css/zt/zTreeStyle.css"/>
	<style>
		#curriculumInfo-divcon {
			width: 0;
			height: 95%;
			background: #fff;
			/*position: absolute;*/
			top: 35px;;
			right: 0;
			/*display:none;*/
			box-shadow: #a9a9a9 0px 0px 10px;
			-webkit-box-shadow: #a9a9a9 0px 0px 10px;
			-moz-box-shadow: #a9a9a9 0px 0px 10px;
			overflow: auto;
			position: fixed;
			z-index: 999;
		}

		#curriculumInfo-divcon1, #curriculumInfo-divcon2 {
			width: 0;
			height: 95%;
			background: #fff;
			/*position: absolute;*/
			top: 35px;;
			right: 0;
			/*display:none;*/
			box-shadow: #a9a9a9 0px 0px 10px;
			-webkit-box-shadow: #a9a9a9 0px 0px 10px;
			-moz-box-shadow: #a9a9a9 0px 0px 10px;
			overflow: auto;
			position: fixed;
			z-index: 999;
		}

		.div-title {
			width: 100%;
			position: relative;
			border-bottom: #eeece8 solid 1px;
		}

		.div-title h3 {
			margin: 0;
			padding: 0;
			font-size: 16px;
			line-height: 55px;
			color: #999;
			text-indent: 23px;
			font-weight: normal;
		}

		.div-title span {
			margin: 0;
			padding: 0;
			text-align: center;
			margin-right: 10px;
			display: inline-block;
			position: absolute;
			right: 0px;
			top: 15px;
			width: 50px;
			height: 25px;
			border-left: #dbdbdb solid 1px;
		}

		.xc_01 {
			width: 40px;
			height: 40px;
			box-shadow: 0px 4px 0 #9abc32;
			border-radius: 20px;
			background: #aeca5b url(/assets/images/xc_00.png) no-repeat center;
		}

		.xc_02 {
			width: 40px;
			height: 40px;
			box-shadow: 0px 4px 0 #6fb3e0;
			border-radius: 20px;
			background: #8cc2e6 url(/assets/images/xc_02.png) no-repeat center;
		}

		.xc_03 {
			width: 40px;
			height: 40px;
			box-shadow: 0px 4px 0 #cb6fd7;
			border-radius: 20px;
			background: #d68cdf url(/assets/images/xc_03.png) no-repeat center;
		}

		.xc_04 {
			width: 40px;
			height: 40px;
			box-shadow: 0px 4px 0 #d53f40;
			border-radius: 20px;
			background: #de6566 url(/assets/images/xc_04.png) no-repeat center;
		}

		.margin-right-10 {
			margin-right: 10px;
			margin-bottom: 10px;
			cursor: pointer;
		}

		.line {
			border-bottom: 0px
		}
	</style>

	
	<link rel="stylesheet" href="/css/zt/zTreeStyle.css"/>
	
	<script type="text/javascript" src="/js/zTree/jquery.ztree.core.min.js"></script>
	
	<script type="text/javascript">

		$(function () {

			$("#xs-info div.col-xs-6").each(function (i, v) {
				if (i % 2 == 0) {
					$(v).css("float", "left");
				} else {
					$(v).css("float", "right");
				}

			});

			if (!(bIsIpad || bIsIphoneOs || bIsMidp || bIsUc7 || bIsUc || bIsAndroid || bIsCE || bIsWM)) {
				$('#loading-btn').on(ace.click_event, function () {
					var btn = $(this);
					btn.button('loading');
				});
			}
		});
		var zTreeObj;

		function pyfa(obj) {
			var fajhh;
			var url;
			if (obj == "zx") {
				fajhh = document.getElementById("zx").value;
			} else {
				fajhh = document.getElementById("fx").value;
			}
			url = "/student/rollManagement/project/qm57234NRX/" + fajhh + "/1/detail";
			$.ajax({
				url: url,
				cache: false,
				type: "get",
				dataType: "json",
				success: function (d) {
					if (d.result && d.result == "error") {
						layer.alert(d.msg);
					} else if (d["treeList"] == null) {
						urp.alert("对不起，没有查询到您的主修方案！");
					} else {
						fillFajh(d);
						$("#tit").html(d["title"]);
						$.fn.zTree.init($("#treeDemo"), setting, d["treeList"]);
						$("#expandAllBtn").bind("click", {type: "expandAll"}, expandNode);
						$("#collapseAllBtn").bind("click", {type: "collapseAll"}, expandNode);
						$("#curriculumInfo-divcon2").animate({width: '70%'});
						$('#curriculumInfo-divcon').animate({width: '0px'});
						$('#curriculumInfo-divcon1').animate({width: '0px'});
					}
				},
				error: function (xhr) {
					urp.alert("错误代码[" + xhr.readyState + "-" + xhr.status + "]:获取数据失败！");
				}
			});
		}

		function fillFajh(data) {
			var jhFajhb = data["jhFajhb"];
			var tcont = "";
			tcont += "<div class='widget-box transparent'>";
			tcont += "<div class='widget-header widget-header-small'>";
			tcont += "<h4 class='widget-title smaller grey'>方案计划信息</h4>";
			tcont += "</div></div>";
			tcont += "<div class='self profile-user-info profile-user-info-striped'>";
			tcont += "<div class='profile-info-row'><div class='profile-info-name' >方案名称</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.famc == null ? "" : jhFajhb.famc) + "";
			tcont += "</div></div>";
			tcont += "<div class='profile-info-row'><div class='profile-info-name'>计划名称</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.jhmc == null ? "" : jhFajhb.jhmc) + "";
			tcont += "</div></div>";
			tcont += "<div class='profile-info-row'><div class='profile-info-name'>年级</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.njmc == null ? "" : jhFajhb.njmc) + "";
			tcont += "</div></div>";
			tcont += "<div class='profile-info-row' ><div class='profile-info-name'>院系名称</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.xsm == null ? "" : jhFajhb.xsm) + "";
			tcont += "</div></div>";
			tcont += "<div class='profile-info-row'><div class='profile-info-name'>专业名称</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.zym == null ? "" : jhFajhb.zym) + "";
			tcont += "</div></div>	";
			tcont += "<div class='profile-info-row' ><div class='profile-info-name'>专业方向名称</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.zyfxm == null ? "" : jhFajhb.zyfxm) + "";
			tcont += "</div></div>";
			tcont += "<div class='profile-info-row'><div class='profile-info-name'>学位</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.xwm == null ? "" : jhFajhb.xwm) + "";
			tcont += "</div></div>	";
			tcont += "<div class='profile-info-row' ><div class='profile-info-name'>毕业类型</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.bylxmc == null ? "" : jhFajhb.bylxmc) + "";
			tcont += "</div></div>";
			tcont += "<div class='profile-info-row'><div class='profile-info-name'>学制类型</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.xzlxmc == null ? "" : jhFajhb.xzlxmc) + "";
			tcont += "</div></div>";
			tcont += "<div class='profile-info-row'><div class='profile-info-name'>修读类型</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.xdlxmc == null ? "" : jhFajhb.xdlxmc) + "";
			tcont += "</div></div>";
			tcont += "<div class='profile-info-row'><div class='profile-info-name'>方案计划类型</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.fajhlx == null ? "" : jhFajhb.fajhlx) + "";
			tcont += "</div></div>";
			tcont += "<div class='profile-info-row' ><div class='profile-info-name'>开始学年代码</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.xnmc == null ? "" : jhFajhb.xnmc) + "";
			tcont += "</div></div>";
			tcont += "<div class='profile-info-row'><div class='profile-info-name'>学期类型代码</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.xqlxm == null ? "" : jhFajhb.xqlxm) + "";
			tcont += "</div></div>	";
			tcont += "<div class='profile-info-row' ><div class='profile-info-name'>开始学期代码</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.xqm == null ? "" : jhFajhb.xqm) + "";
			tcont += "</div></div>";
			tcont += "<div class='profile-info-row'><div class='profile-info-name'>要求总学分</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.yqzxf == null ? "" : jhFajhb.yqzxf) + "";
			tcont += "</div></div>	";
			tcont += "<div class='profile-info-row'><div class='profile-info-name'>课程总学分</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.kczxf == null ? "" : jhFajhb.kczxf) + "";
			tcont += "</div></div>";
			tcont += "<div class='profile-info-row'><div class='profile-info-name'>课程总门数</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.kczms == null ? "" : jhFajhb.kczms) + "";
			tcont += "</div></div>	";
			tcont += "<div class='profile-info-row' ><div class='profile-info-name'>课程总学时</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.kczxs == null ? "" : jhFajhb.kczxs) + "";
			tcont += "</div></div>";
			tcont += "<div class='profile-info-row'><div class='profile-info-name'>培养目标</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.pymb == null ? "" : jhFajhb.pymb) + "";
			tcont += "</div></div>";
			tcont += "<div class='profile-info-row'><div class='profile-info-name'>修读要求</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.xdyq == null ? "" : jhFajhb.xdyq) + "";
			tcont += "</div></div>";
			tcont += "<div class='profile-info-row'><div class='profile-info-name'>备注</div>";
			tcont += "<div class='profile-info-value'>" + (jhFajhb.bz == null ? "" : jhFajhb.bz) + "";
			tcont += "</div></div></div>";
			$("#fajh").html(tcont);
		}

		var setting = {
			data: {
				simpleData: {
					enable: true
				}
			},
			view: {
				showTitle: false,
				nameIsHTML: true,
				fontCss: getFontCss,
				showIcon: false
			},
			callback: {
				beforeClick: beforeClick
			}
		};

		function getFontCss(treeId, treeNode) {
			return (!!treeNode.highlight) ? {color: "#A60000", "font-weight": "bold"} : {
				color: "#333",
				"font-weight": "normal"
			};
		}

		function expandNode(e) {
			var zTree = $.fn.zTree.getZTreeObj("treeDemo"),
					type = e.data.type,
					nodes = zTree.getSelectedNodes();
			if (type == "expandAll") {
				zTree.expandAll(true);
			} else if (type == "collapseAll") {
				zTree.expandAll(false);
			}
		}

		function beforeClick(treeId, treeNode, clickFlag) {
			//$(".col-xs-6").hide();
			$("#kz").hide();
			$("#fajh").hide();
			$("#kc").hide();
			$("#kcfa").hide();
			$("#xnxq").hide();
			$.get(treeNode.urlPath, function (data) {
				if (data.flag == '1') {
					$("#kz_kzlx").text(data.kz.kzlx == null ? '' : data.kz.kzlx);
					$("#kz_kzm").text(data.kz.kzm == null ? '' : data.kz.kzm);
					$("#kz_kczxf").text(data.kz.kczxf == null ? '' : data.kz.kczxf);
					$("#kz_kczxs").text(data.kz.kczxs == null ? '' : data.kz.kczxs);
					$("#kz_zsms").text(data.kz.zsms == null ? '' : data.kz.zsms);
					$("#kz_zsxf").text(data.kz.zsxf == null ? '' : data.kz.zsxf);
					$("#kz_zsxs").text(data.kz.zsxs == null ? '' : data.kz.zsxs);
					$("#kz_scyq").text(data.kz.scyq == null ? '' : data.kz.scyq);
					$("#kz_bz").text(data.kz.bz == null ? '' : data.kz.bz);
					$("#kz").show();
				}
				if (data.flag == '2') {
					$("#kc_kch").text(data.kc.kch == null ? '' : data.kc.kch);
					$("#kc_kcm").text(data.kc.kcm == null ? '' : data.kc.kcm);
					$("#kc_ywkcm").text(data.kc.ywkcm == null ? '' : data.kc.ywkcm);
					$("#kc_xsm").text(data.kc.xsm == null ? '' : data.kc.xsm);
					$("#kc_kkxq").text(data.kc.kkxq == null ? '' : data.kc.kkxq);
					$("#kc_bybz").text(data.kc.bybz == null ? '' : data.kc.bybz);
					$("#kc_xf").text(data.kc.xf == null ? '' : fmoney(data.kc.xf, 1));
					$("#kc_xs").text(data.kc.xs == null ? '' : data.kc.xs);
					$("#kc_ksrq").text(data.kc.ksrq == null ? '' : data.kc.ksrq);
					$("#kc_jsrq").text(data.kc.jsrq == null ? '' : data.kc.jsrq);
					$("#kc_kcztmc").text(data.kc.kcztsm == null ? '' : data.kc.kcztsm);
					$("#kc_knzxs").text(data.kc.knzxs == null ? '' : data.kc.knzxs);
					$("#kc_jkzxs").text(data.kc.jkzxs == null ? '' : data.kc.jkzxs);
					$("#kc_sjzxs").text(data.kc.sjzxs == null ? '' : data.kc.sjzxs);
					$("#kc_syzxs").text(data.kc.syzxs == null ? '' : data.kc.syzxs);
					$("#kc_qzsjzxs").text(data.kc.qzsjzxs == null ? '' : data.kc.qzsjzxs);
					$("#kc_tlfdzxs").text(data.kc.tlfdzxs == null ? '' : data.kc.tlfdzxs);
					$("#kc_sjzyzxs").text(data.kc.sjzyzxs == null ? '' : data.kc.sjzyzxs);
					$("#kc_kwzxs").text(data.kc.kwzxs == null ? '' : data.kc.kwzxs);
					$("#kc_kwxf").text(data.kc.kwxf == null ? '' : data.kc.kwxf);
					$("#kc_kclbmc").text(data.kc.kclbmc == null ? '' : data.kc.kclbmc);
					$("#kc_jxfssm").text(data.kc.jxfssm == null ? '' : data.kc.jxfssm);
					$("#kc_jsm").text(data.kc.jsm == null ? '' : data.kc.jsm);
					$("#kc_jc").text(data.kc.jc == null ? '' : data.kc.jc);
					$("#kc_cks").text(data.kc.cks == null ? '' : data.kc.cks);
					$("#kc_szdw").text(data.kc.szdw == null ? '' : data.kc.szdw);
					$("#kc_kcsm").text(data.kc.kcsm == null ? '' : data.kc.kcsm);
					$("#kc_kslxmc").text(data.kc.kslxmc == null ? '' : data.kc.kslxmc);
					$("#kc_xaqm").text(data.kc.xaqm == null ? '' : data.kc.xaqm);
					$("#kc_nrjj").text(data.kc.nrjj == null ? '' : data.kc.nrjj);
					$("#kc_bz").text(data.kc.bz == null ? '' : data.kc.bz);
					$("#kcfa_xnmc").text(data.jhkc.xnmc == null ? '' : data.jhkc.xnmc);
					$("#kcfa_xqm").text(data.jhkc.xqm == null ? '' : data.jhkc.xqm);
					$("#kcfa_kcsxmc").text(data.jhkc.kcsxmc == null ? '' : data.jhkc.kcsxmc);
					$("#kcfa_kslxmc").text(data.jhkc.kslxmc == null ? '' : data.jhkc.kslxmc);
					$("#kcfa_bz").text(data.jhkc.bz == null ? '' : data.jhkc.bz);
					$("#kcfa_bz1").text(data.jhkc.bz1 == null ? '' : data.jhkc.bz1);
					$("#kcfa_bz2").text(data.jhkc.bz2 == null ? '' : data.jhkc.bz2);
					$("#kcfa_bz3").text(data.jhkc.bz3 == null ? '' : data.jhkc.bz3);
					$("#kc").show();
					$("#kcfa").show();
				}
				if (data.flag == '3') {
					$("#xnxq_jhxn").text(data.jhxnxq.id.jhxn == null ? '' : data.jhxnxq.id.jhxn);
					$("#xnxq_xqlx").text(data.jhxnxq.id.xqlxdm == null ? '' : data.jhxnxq.id.xqlxdm);
					$("#xnxq_jhxq").text(data.jhxnxq.id.xqm == null ? '' : data.jhxnxq.id.xqm);
					$("#xnxq_zsxdms").text(data.jhxnxq.zsms == null ? '' : data.jhxnxq.zsms);
					$("#xnxq_zdxdms").text(data.jhxnxq.zdms == null ? '' : data.jhxnxq.zdms);
					$("#xnxq_zsxdxf").text(data.jhxnxq.zsxf == null ? '' : data.jhxnxq.zsxf);
					$("#xnxq_zdxdxf").text(data.jhxnxq.zdxf == null ? '' : data.jhxnxq.zdxf);
					$("#xnxq_zsxdxs").text(data.jhxnxq.zsxs == null ? '' : data.jhxnxq.zsxs);
					$("#xnxq_zdxdxs").text(data.jhxnxq.zdxs == null ? '' : data.jhxnxq.zdxs);
					$("#xnxq_scyq").text(data.jhxnxq.scyq == null ? '' : data.jhxnxq.scyq);
					$("#xnxq_bz").text(data.jhxnxq.bz == null ? '' : data.jhxnxq.bz);
					$("#xnxq").show();
				}
				if (data.flag == '4') {
					$("#xnxq_jhxn").text(data.jhxnxq.id.jhxn == null ? '' : data.jhxnxq.id.jhxn);
					$("#xnxq_xqlx").text(data.jhxnxq.id.xqlxdm == null ? '' : data.jhxnxq.id.xqlxdm);
					$("#xnxq_jhxq").text(data.jhxnxq.id.xqm == null ? '' : data.jhxnxq.id.xqm);
					$("#xnxq_zsxdms").text(data.jhxnxq.zsms == null ? '' : data.jhxnxq.zsms);
					$("#xnxq_zdxdms").text(data.jhxnxq.zdms == null ? '' : data.jhxnxq.zdms);
					$("#xnxq_zsxdxf").text(data.jhxnxq.zsxf == null ? '' : data.jhxnxq.zsxf);
					$("#xnxq_zdxdxf").text(data.jhxnxq.zdxf == null ? '' : data.jhxnxq.zdxf);
					$("#xnxq_zsxdxs").text(data.jhxnxq.zsxs == null ? '' : data.jhxnxq.zsxs);
					$("#xnxq_zdxdxs").text(data.jhxnxq.zdxs == null ? '' : data.jhxnxq.zdxs);
					$("#xnxq_scyq").text(data.jhxnxq.scyq == null ? '' : data.jhxnxq.scyq);
					$("#xnxq_bz").text(data.jhxnxq.bz == null ? '' : data.jhxnxq.bz);
					$("#xnxq").show();
				}
				if (data.flag == '5') {
					$("#kc_kch").text(data.kc.kch == null ? '' : data.kc.kch);
					$("#kc_kcm").text(data.kc.kcm == null ? '' : data.kc.kcm);
					$("#kc_ywkcm").text(data.kc.ywkcm == null ? '' : data.kc.ywkcm);
					$("#kc_xsm").text(data.kc.xsm == null ? '' : data.kc.xsm);
					$("#kc_kkxq").text(data.kc.kkxq == null ? '' : data.kc.kkxq);
					$("#kc_bybz").text(data.kc.bybz == null ? '' : data.kc.bybz);
					$("#kc_xf").text(data.kc.xf == null ? '' : data.kc.xf);
					$("#kc_xs").text(data.kc.xs == null ? '' : data.kc.xs);
					$("#kc_ksrq").text(data.kc.ksrq == null ? '' : data.kc.ksrq);
					$("#kc_jsrq").text(data.kc.jsrq == null ? '' : data.kc.jsrq);
					$("#kc_kcztmc").text(data.kc.kcztsm == null ? '' : data.kc.kcztsm);
					$("#kc_knzxs").text(data.kc.knzxs == null ? '' : data.kc.knzxs);
					$("#kc_jkzxs").text(data.kc.jkzxs == null ? '' : data.kc.jkzxs);
					$("#kc_sjzxs").text(data.kc.sjzxs == null ? '' : data.kc.sjzxs);
					$("#kc_syzxs").text(data.kc.syzxs == null ? '' : data.kc.syzxs);
					$("#kc_qzsjzxs").text(data.kc.qzsjzxs == null ? '' : data.kc.qzsjzxs);
					$("#kc_tlfdzxs").text(data.kc.tlfdzxs == null ? '' : data.kc.tlfdzxs);
					$("#kc_sjzyzxs").text(data.kc.sjzyzxs == null ? '' : data.kc.sjzyzxs);
					$("#kc_kwzxs").text(data.kc.kwzxs == null ? '' : data.kc.kwzxs);
					$("#kc_kwxf").text(data.kc.kwxf == null ? '' : data.kc.kwxf);
					$("#kc_kclbmc").text(data.kc.kclbmc == null ? '' : data.kc.kclbmc);
					$("#kc_jxfssm").text(data.kc.jxfssm == null ? '' : data.kc.jxfssm);
					$("#kc_jsm").text(data.kc.jsm == null ? '' : data.kc.jsm);
					$("#kc_jc").text(data.kc.jc == null ? '' : data.kc.jc);
					$("#kc_cks").text(data.kc.cks == null ? '' : data.kc.cks);
					$("#kc_szdw").text(data.kc.szdw == null ? '' : data.kc.szdw);
					$("#kc_kcsm").text(data.kc.kcsm == null ? '' : data.kc.kcsm);
					$("#kc_kslxmc").text(data.kc.kslxmc == null ? '' : data.kc.kslxmc);
					$("#kc_xqm").text(data.kc.xqm == null ? '' : data.kc.xqm);
					$("#kc_nrjj").text(data.kc.nrjj == null ? '' : data.kc.nrjj);
					$("#kc_bz").text(data.kc.bz == null ? '' : data.kc.bz);
					$("#kc").show();
				}
			});
			return (treeNode.click != false);
		}

		function fmoney(s, n) {
			/*
			 * 参数说明：
			 * s：要格式化的数字
			 * n：保留几位小数
			 * */
			n = n > 0 && n <= 20 ? n : 2;
			s = parseFloat((s + "").replace(/[^\d\.-]/g, "")).toFixed(n) + "";
			var l = s.split(".")[0].split("").reverse(),
					r = s.split(".")[1];
			t = "";
			for (i = 0; i < l.length; i++) {
				t += l[i] + ((i + 1) % 3 == 0 && (i + 1) != l.length ? "," : "");
			}
			return t.split("").reverse().join("") + "." + r;
		}

		/* $(function(){
		 $.fn.zTree.init($("#treeDemo"), setting, zNodes);
		 $("#expandAllBtn").bind("click", {type:"expandAll"}, expandNode);
		 $("#collapseAllBtn").bind("click", {type:"collapseAll"}, expandNode);
		 }); */

		function xjydjc(f) {
			$.ajax({
				url: "/student/rollManagement/project/queryInfoById",
				cache: false,
				type: "get",
				dataType: "json",
				success: function (d) {
					$("#curriculumInfo-divcon1").animate({width: '70%'});
					$('#curriculumInfo-divcon').animate({width: '0px'});
					$('#curriculumInfo-divcon2').animate({width: '0px'});
				},
				error: function (xhr) {
					$("#curriculumInfo-divcon1").animate({width: '70%'});
					$('#curriculumInfo-divcon').animate({width: '0px'});
					$('#curriculumInfo-divcon2').animate({width: '0px'});
				}
			});
		}

		function xjydjc1(f) {
			$.ajax({
				url: "/student/rollManagement/project/queryInfoById",
				cache: false,
				type: "get",
				dataType: "json",
				success: function (d) {
					$("#curriculumInfo-divcon").animate({width: '70%'});

					$('#curriculumInfo-divcon1').animate({width: '0px'});
					$('#curriculumInfo-divcon2').animate({width: '0px'});
				},
				error: function (xhr) {
					$("#curriculumInfo-divcon").animate({width: '70%'});
					$('#curriculumInfo-divcon1').animate({width: '0px'});
					$('#curriculumInfo-divcon2').animate({width: '0px'});
				}
			});
		}

		$(function () {
			$("#curriculumInfo-divcon").click(function () {
				return false;
			});
			$("#curriculumInfo-divcon1").click(function () {
				return false;
			});
			$("#curriculumInfo-divcon2").click(function () {
				return false;
			});

			$(document).click(function (e) {
				var id = $(e.target).attr("id");
				if (id == "td_div" || $(e.target).closest("#td_div").size() > 0) {

				} else {
					$('#curriculumInfo-divcon').animate({width: '0px'});
					$('#curriculumInfo-divcon1').animate({width: '0px'});
					$('#curriculumInfo-divcon2').animate({width: '0px'});
				}
			});
		});


		function xiugai() {
			location.href = "/student/rollManagement/personalInfoUpdate/index";
		}

		function tishi() {

		}
		
		let map = new Map();
		map.set("email", "qq3392313023@163.com");
		map.set("dh", "17353878918");
		map.set("brlxdh", "");
		map.set("sfzh", "370112200407117114");
		map.set("rxrq", "");
		
		function showInfo(flag, obj) {
			$(obj).html(map.get(flag));
			$(obj).attr("onclick", "");
			$(obj).attr("style", "");
		}
	</script>



</head>
<body class="no-skin">

<!-- 查看详情专用 -->
<div class="modal fade" id="view-table" tabindex="1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
	<div class="modal-dialog" style="width:60%;">
		<div class="modal-content">
			<div class="center">
				<span class="view-pre-loading"></span>
			</div>
		</div>
	</div>
</div>
<!-- 查看详情专用 -->

<div class="modal fade" id="myModal" tabindex="1" role="dialog"
	 aria-labelledby="myModalLabel" aria-hidden="true">
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
				<button type="button" class="close" data-dismiss="modal"
						aria-hidden="true">×
				</button>
				<h4 class="modal-title" id="myModalLabel">
					调课后的周次
				</h4>
			</div>
			<div class="modal-body">
				开始周次:<input type="text" name="kszc" id="kszc" readonly="readonly"/><br/>
				结束周次:<input type="text" name="jszc" id="jszc"/><br/>
				选择周次模式:<input type="radio" name="zcms" value="1" style="width:20px"/> 全周 &nbsp;&nbsp;
				<input type="radio" name="zcms" value="2" style="width:20px"/> 单周&nbsp;&nbsp;
				<input type="radio" name="zcms" value="3" style="width:20px"/> 双周
			</div>
			<div class="modal-footer">
				<button type="button" class="btn btn-default">
					保存
				</button>
				<button type="button" class="btn btn-primary" data-dismiss="modal">
					关闭
				</button>
			</div>
		</div>
		<!-- /.modal-content -->
	</div>
	<!-- /.modal-dialog -->
</div>
<div id="navbar" class="navbar navbar-default navbar-fixed-top">
	<script type="text/javascript">
		try {
			ace.settings.check('navbar', 'fixed')
		} catch (e) {
		}
	</script>

	<div class="navbar-container" id="navbar-container">
		<button type="button" class="navbar-toggle menu-toggler pull-left" id="menu-toggler">
			<span class="sr-only">Toggle sidebar</span>

			<span class="icon-bar"></span>

			<span class="icon-bar"></span>

			<span class="icon-bar"></span>
		</button>
		<div class="navbar-header pull-left">
			<a href="/" class="navbar-brand" id='0' onclick="toSelect(this);">
				<small style="font-weight:700;font-family:微软雅黑">
					<!--
					<i class="fa fa-leaf"></i>
					Ace Admin
					 -->
					
					URP高校教务管理系统

				</small>
			</a>
		</div>

		<div class="navbar-buttons navbar-header pull-right" role="navigation">
			<ul class="nav ace-nav">
				<li class="green" style="text-align: center">
					<div class="intellegenceUDiv" style="z-index:1" id="intellegenceUDiv"
						 style="display:inline-block; left: 100%; position:absolute;height:100%;line-height:normal;">
						<div style="height:100%; width:32px; display:inline-block;
									 position:absolute;line-height:100%;left:-32px;
									 background-color:#777; " id="clickdiv"
							 onclick="changeInfo()">
							<i class="fa fa-search white" style='margin-top: 50%;' id="clicki"></i>
						</div>
					</div>
					<div id="form-search" class="nav-search" style="width:0px;right:0px;margin-top: 8px; margin-right:20px;background-color:rgba(0,0,0,0);z-index:-1">
						<form class="form-search">
							<span class="input-icon">
								<input type="text" placeholder="查找功能..." class="nav-search-input" id="search-input"
									   autocomplete="off">
								<i class="ace-icon fa fa-search blue"></i>
							</span>
						</form>
					</div>
					<!-- 通知   -->

						
				<!-- 展示当前的学期季度信息 -->
				<li class="light-red" style="text-align: center">
					<a data-toggle="dropdown" class="dropdown-toggle" href="#" onclick="SchoolCalendar()">
						<i class="ace-icon fa fa-calendar"></i>


2024-2025 秋  第11周   星期五
					</a>
				</li>
				

				
				<li class="light-blue" style="text-align: center">
					<a data-toggle="dropdown" href="#" class="dropdown-toggle">
						<img class="nav-user-photo" src="/img/head/man.png" onerror="this.src='/img/head/man.png'"
							 alt="Jason's Photo"/>
						<span class="user-info">
							<small>欢迎您，</small>
							


张贤德
						</span>
						<i class="ace-icon fa fa-caret-down"></i>
					</a>

					<ul class="user-menu dropdown-menu-right dropdown-menu dropdown-yellow dropdown-caret dropdown-close">
						<li>
							<a href="/" title="主页" id='0' onclick="toSelect(this);">
								<i class="ace-icon fa fa-home"></i>首页
							</a>
						</li>

						<li>
							<a href="javascript:changePassword('/student/rollManagement/personalInfoUpdate/updatePassword')">
								<i class="ace-icon fa fa-user"></i>
								修改密码
							</a>
						</li>

						<li class="divider"></li>

						<li>
							<a id="logout" href="/logout" title="退出系统">
								<i class="ace-icon fa fa-power-off"></i>
								注销
							</a>
						</li>
					</ul>
				</li>
				
				<!-- 主页 -->
				<!-- <li class="light-blue" style="text-align: center" >
					 <a href="/" title="主页" id='0' onclick="toSelect(this);">
						<i class="ace-icon fa fa-home"></i>
					</a>
				</li>
				<li  class="light-blue" style="text-align: center" >
					<a href="/logout" title="退出系统" onclick="removeCookie();">
						<i class="ace-icon fa fa-power-off"></i>
					</a>
				</li>
-->
			</ul>
		</div>
	</div>
	<!-- /.navbar-container -->
</div>

<div class="main-container" id="main-container">
	<script type="text/javascript">
		function changeInfo() {
			$("#form-search").animate({width: '152px', right: '22px'});
			//margin-top: 8px; margin-right:20px;
		}

		$(document).click(function (e) {
			var id = $(e.target).attr("id");
			if (id == "clickdiv" || id == "clicki" || id == "search-input") {

			} else {
				$("#form-search").animate({width: '0px', right: '0px'});
			}
		})
		try {
			ace.settings.check('main-container', 'fixed')
		} catch (e) {
		}

		function SchoolCalendar() {
			//	var ifheight = $(window).height()-$(".modal-content").offset().top ;
			$(".modal-content").html("<iframe width='100%' style='width:100%; min-height:600px; height:100%; overflow: hidden;' src='/indexCalendar' ></iframe>");
			//$(".modal-content iframe").css("max-height",ifheight);
			$("#view-table").modal().show().on("hidden.bs.modal", function () {
				$(this).removeData("bs.modal");
			});
		}

		var sUserAgent = navigator.userAgent.toLowerCase();
		var bIsIpad = sUserAgent.match(/ipad/i) == "ipad";
		var bIsIphoneOs = sUserAgent.match(/iphone os/i) == "iphone os";
		var bIsMidp = sUserAgent.match(/midp/i) == "midp";
		var bIsUc7 = sUserAgent.match(/rv:1.2.3.4/i) == "rv:1.2.3.4";
		var bIsUc = sUserAgent.match(/ucweb/i) == "ucweb";
		var bIsAndroid = sUserAgent.match(/android/i) == "android";
		var bIsCE = sUserAgent.match(/windows ce/i) == "windows ce";
		var bIsWM = sUserAgent.match(/windows mobile/i) == "windows mobile";

		var isMenuToDeal = true;

		function rootMenuClick(obj) {
			if (isMenuToDeal && !$(obj).hasClass("open")) {
				$.each($(obj).find(".hsub"), function () {
					if ($(this).hasClass("open")) {
						isMenuToDeal = false;
					}
				});
			} else {
				isMenuToDeal = false;
			}

			if (isMenuToDeal) {
				var firobj = $(obj).children("ul").children("li").first();
				firobj.addClass("open");
				$(firobj).children("ul").css("display", "block");
			} else {
				isMenuToDeal = true;
			}
		}

		function stopHere(e) {
			isMenuToDeal = false;
		}

		/*菜单上的搜索框*/
		$(document).ready(function () {
			var obj_lis = $("#menus").find("li > ul > li > ul > li");
			var obj_menus = new Array();
			for (i = 0; i < obj_lis.length; i++) {
				obj_menus[i] = $.trim($(obj_lis[i]).find("a:first").text());
			}
			try {
				$("#search-input").bs_typeahead({
					source: obj_menus,
					updater: function (a) {
						for (i = 0; i < obj_lis.length; i++) {
							if (a == $.trim($(obj_lis[i]).find("a:first").text())) {
								$.cookie('selectionBar', obj_lis[i].id, {
									path: '/'
								});
								window.location.href = $(obj_lis[i]).find("a:first").attr("href");
								break;
							}
						}
						return a;
					}
				})
			} catch (a) {
			}
		});

		function changePassword(url) {
			///student/rollManagement/personalInfoUpdate/updatePassword
			$('#view-table').modal({
				remote: url
			}).on("hidden.bs.modal", function () {
				/*$("div").remove(".widget-box");*/
				$(this).removeData("bs.modal");
			}).css({
				width: 'auto'
			});
		}

	</script>

	<style type="text/css">
		.sidebar > .nav-search {
			position: static;
			background-color: #FAFAFA;
			border-bottom: 1px solid #DDD;
			text-align: center;
			height: 35px;
			padding-top: 0px;
		}

		.sidebar > .nav-search .nav-search-input {
			width: 162px !important;
			border-radius: 5px !important;
			max-width: 162px !important;
			opacity: 1 !important;
			filter: alpha(opacity=100) !important;
		}

		#breadcrumbs li {
			cursor: pointer;
		}
	</style>
	<div id="sidebar" class="sidebar responsive sidebar-fixed">
		<script type="text/javascript">
			try {
				ace.settings.check('sidebar', 'fixed')
			} catch (e) {
			}
		</script>

		<ul class="nav nav-list" id="menus">
			

			
			<li class="" id="1181690" onclick="rootMenuClick(this);">
				<a href="#" class="dropdown-toggle">
					<i class="menu-icon fa fa-user"></i>
					<span class="menu-text">个人管理
					</span>
					<b class="arrow fa fa-angle-down"></b>
				</a>

				<b class="arrow"></b>
				
				<ul class="submenu" onclick="stopHere();">
					
					<li class="" id="1188620">
						<a href="#" class="dropdown-toggle">
							<i class="menu-icon fa fa-caret-right"></i>
							学籍管理
							<b class="arrow fa fa-angle-down"></b>
						</a>

						<b class="arrow"></b>

						
						<ul class="submenu">
							
							<li class="" id="1183421" onclick="toSelect(this);">
								<a href="/student/rollManagement/rollInfo/index">&nbsp;&nbsp;
									学生学籍信息
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1183422" onclick="toSelect(this);">
								<a href="/student/rollManagement/personalInfoUpdate/index">&nbsp;&nbsp;
									个人信息修改
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1183423" onclick="toSelect(this);">
								<a href="/student/rollManagement/rollChanges/index">&nbsp;&nbsp;
									学籍异动
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1183424" onclick="toSelect(this);">
								<a href="/student/rollManagement/rewardsAndPenalties/index">&nbsp;&nbsp;
									奖惩信息
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1183425" onclick="toSelect(this);">
								<a href="/student/rollManagement/electronicRegistration/index">&nbsp;&nbsp;
									电子注册
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1183426" onclick="toSelect(this);">
								<a href="/student/rollManagement/minorProgramRegistration/index">&nbsp;&nbsp;
									辅修方案注册
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="125803404" onclick="toSelect(this);">
								<a href="/student/rollManagement/informationCollection/index">&nbsp;&nbsp;
									监护人信息采集
								</a>

								<b class="arrow"></b>
							</li>
							
						</ul>

						
					</li>
					
					<li class="" id="1188621">
						<a href="#" class="dropdown-toggle">
							<i class="menu-icon fa fa-caret-right"></i>
							学生异动
							<b class="arrow fa fa-angle-down"></b>
						</a>

						<b class="arrow"></b>

						
						<ul class="submenu">
							
							<li class="" id="1183428" onclick="toSelect(this);">
								<a href="/student/personalManagement/studentChange/index">&nbsp;&nbsp;
									学籍异动申请
								</a>

								<b class="arrow"></b>
							</li>
							
						</ul>

						
					</li>
					
					<li class="" id="1188622">
						<a href="#" class="dropdown-toggle">
							<i class="menu-icon fa fa-caret-right"></i>
							毕业设计
							<b class="arrow fa fa-angle-down"></b>
						</a>

						<b class="arrow"></b>

						
						<ul class="submenu">
							
							<li class="" id="1183620" onclick="toSelect(this);">
								<a href="/student/personalManagement/projectSelect/index">&nbsp;&nbsp;
									网上选题
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1183621" onclick="toSelect(this);">
								<a href="/student/personalManagement/paperSubmit/index">&nbsp;&nbsp;
									论文提交
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1183622" onclick="toSelect(this);">
								<a href="/student/personalManagement/projectScore/index">&nbsp;&nbsp;
									毕业设计成绩查询
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1183623" onclick="toSelect(this);">
								<a href="/student/personalManagement/excellentProject/index">&nbsp;&nbsp;
									优秀毕业设计名单查询
								</a>

								<b class="arrow"></b>
							</li>
							
						</ul>

						
					</li>
					
				</ul>
				
			</li>
			
			<li class="" id="82020" onclick="rootMenuClick(this);">
				<a href="#" class="dropdown-toggle">
					<i class="menu-icon fa fa-shopping-cart"></i>
					<span class="menu-text">选课管理
					</span>
					<b class="arrow fa fa-angle-down"></b>
				</a>

				<b class="arrow"></b>
				
				<ul class="submenu" onclick="stopHere();">
					
					<li class="" id="82021">
						<a href="#" class="dropdown-toggle">
							<i class="menu-icon fa fa-caret-right"></i>
							本学期课表
							<b class="arrow fa fa-angle-down"></b>
						</a>

						<b class="arrow"></b>

						
						<ul class="submenu">
							
							<li class="" id="82022" onclick="toSelect(this);">
								<a href="/student/courseSelect/thisSemesterCurriculum/index">&nbsp;&nbsp;
									本学期课表
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1162620" onclick="toSelect(this);">
								<a href="/student/courseSelect/calendarSemesterCurriculum/index">&nbsp;&nbsp;
									历年学期课表
								</a>

								<b class="arrow"></b>
							</li>
							
						</ul>

						
					</li>
					
					<li class="" id="1293220">
						<a href="#" class="dropdown-toggle">
							<i class="menu-icon fa fa-caret-right"></i>
							选课管理
							<b class="arrow fa fa-angle-down"></b>
						</a>

						<b class="arrow"></b>

						
						<ul class="submenu">
							
							<li class="" id="1293217" onclick="toSelect(this);">
								<a href="/student/courseSelect/courseSelectNotice/index">&nbsp;&nbsp;
									选课公告
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1293218" onclick="toSelect(this);">
								<a href="/student/courseSelect/courseSelect/index">&nbsp;&nbsp;
									选课
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1293219" onclick="toSelect(this);">
								<a href="/student/courseSelect/courseSelectResult/index">&nbsp;&nbsp;
									选课结果
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1293221" onclick="toSelect(this);">
								<a href="/student/courseSelect/quitCourse/index">&nbsp;&nbsp;
									退课
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1298420" onclick="toSelect(this);">
								<a href="/student/courseSelect/courseSelectFailed/index">&nbsp;&nbsp;
									选课失败信息
								</a>

								<b class="arrow"></b>
							</li>
							
						</ul>

						
					</li>
					
				</ul>
				
			</li>
			
			<li class="" id="12580300" onclick="rootMenuClick(this);">
				<a href="#" class="dropdown-toggle">
					<i class="menu-icon fa fa-pencil"></i>
					<span class="menu-text">教学评估
					</span>
					<b class="arrow fa fa-angle-down"></b>
				</a>

				<b class="arrow"></b>
				
				<ul class="submenu" onclick="stopHere();">
					
					<li class="" id="12580301">
						<a href="#" class="dropdown-toggle">
							<i class="menu-icon fa fa-caret-right"></i>
							评估
							<b class="arrow fa fa-angle-down"></b>
						</a>

						<b class="arrow"></b>

						
						<ul class="submenu">
							
							<li class="" id="12580302" onclick="toSelect(this);">
								<a href="/student/teachingEvaluation/evaluation/index">&nbsp;&nbsp;
									教学评估
								</a>

								<b class="arrow"></b>
							</li>
							
						</ul>

						
					</li>
					
				</ul>
				
			</li>
			
			<li class="" id="1423377" onclick="rootMenuClick(this);">
				<a href="#" class="dropdown-toggle">
					<i class="menu-icon fa fa-list-alt"></i>
					<span class="menu-text">考务管理
					</span>
					<b class="arrow fa fa-angle-down"></b>
				</a>

				<b class="arrow"></b>
				
				<ul class="submenu" onclick="stopHere();">
					
					<li class="" id="1423378">
						<a href="#" class="dropdown-toggle">
							<i class="menu-icon fa fa-caret-right"></i>
							考务管理
							<b class="arrow fa fa-angle-down"></b>
						</a>

						<b class="arrow"></b>

						
						<ul class="submenu">
							
							<li class="" id="1423379" onclick="toSelect(this);">
								<a href="/student/examinationManagement/examPlan/index">&nbsp;&nbsp;
									考试安排
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1423380" onclick="toSelect(this);">
								<a href="/student/examinationManagement/examSignUp/index">&nbsp;&nbsp;
									考试报名
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1423381" onclick="toSelect(this);">
								<a href="/student/examinationManagement/examGrade/index">&nbsp;&nbsp;
									考试成绩
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1423382" onclick="toSelect(this);">
								<a href="/student/examinationManagement/cetSignUp/index">&nbsp;&nbsp;
									四六级报名
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1423383" onclick="toSelect(this);">
								<a href="/student/examinationManagement/cetGrade/index">&nbsp;&nbsp;
									四六级成绩
								</a>

								<b class="arrow"></b>
							</li>
							
						</ul>

						
					</li>
					
				</ul>
				
			</li>
			
			<li class="" id="1423370" onclick="rootMenuClick(this);">
				<a href="#" class="dropdown-toggle">
					<i class="menu-icon fa fa-building-o"></i>
					<span class="menu-text">教学资源
					</span>
					<b class="arrow fa fa-angle-down"></b>
				</a>

				<b class="arrow"></b>
				
				<ul class="submenu" onclick="stopHere();">
					
					<li class="" id="1423371">
						<a href="#" class="dropdown-toggle">
							<i class="menu-icon fa fa-caret-right"></i>
							教学资源
							<b class="arrow fa fa-angle-down"></b>
						</a>

						<b class="arrow"></b>

						
						<ul class="submenu">
							
							<li class="" id="1423372" onclick="toSelect(this);">
								<a href="/student/teachingResources/classroomCurriculum/index">&nbsp;&nbsp;
									教室课表
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1423373" onclick="toSelect(this);">
								<a href="/student/teachingResources/teacherCurriculum/index">&nbsp;&nbsp;
									教师课表
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1423374" onclick="toSelect(this);">
								<a href="/student/teachingResources/classCurriculum/index">&nbsp;&nbsp;
									班级课表
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1423384" onclick="toSelect(this);">
								<a href="/student/teachingResources/courseCurriculum/index">&nbsp;&nbsp;
									课程课表
								</a>

								<b class="arrow"></b>
							</li>
							
						</ul>

						
					</li>
					
					<li class="" id="1443371">
						<a href="#" class="dropdown-toggle">
							<i class="menu-icon fa fa-caret-right"></i>
							自习查询
							<b class="arrow fa fa-angle-down"></b>
						</a>

						<b class="arrow"></b>

						
						<ul class="submenu">
							
							<li class="" id="1443372" onclick="toSelect(this);">
								<a href="/student/teachingResources/freeClassroom/index">&nbsp;&nbsp;
									空闲教室查询
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1443384" onclick="toSelect(this);">
								<a href="/student/teachingResources/classroomUseStatus/index">&nbsp;&nbsp;
									教室使用状况查询
								</a>

								<b class="arrow"></b>
							</li>
							
						</ul>

						
					</li>
					
				</ul>
				
			</li>
			
			<li class="" id="1377620" onclick="rootMenuClick(this);">
				<a href="#" class="dropdown-toggle">
					<i class="menu-icon fa fa-search-plus"></i>
					<span class="menu-text">综合查询
					</span>
					<b class="arrow fa fa-angle-down"></b>
				</a>

				<b class="arrow"></b>
				
				<ul class="submenu" onclick="stopHere();">
					
					<li class="" id="1377621">
						<a href="#" class="dropdown-toggle">
							<i class="menu-icon fa fa-caret-right"></i>
							成绩查询
							<b class="arrow fa fa-angle-down"></b>
						</a>

						<b class="arrow"></b>

						
						<ul class="submenu">
							
							<li class="" id="125803405" onclick="toSelect(this);">
								<a href="/student/integratedQuery/scoreQuery/allTermScores/index">&nbsp;&nbsp;
									历年成绩
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1379870" onclick="toSelect(this);">
								<a href="/student/integratedQuery/scoreQuery/allPassingScores/index">&nbsp;&nbsp;
									全部及格成绩
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1417250" onclick="toSelect(this);">
								<a href="/student/integratedQuery/scoreQuery/coursePropertyScores/index">&nbsp;&nbsp;
									课程属性成绩
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1443373" onclick="toSelect(this);">
								<a href="/student/integratedQuery/scoreQuery/schemeScores/index">&nbsp;&nbsp;
									方案成绩
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1443374" onclick="toSelect(this);">
								<a href="/student/integratedQuery/scoreQuery/unpassedScores/index">&nbsp;&nbsp;
									不及格成绩
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1443375" onclick="toSelect(this);">
								<a href="/student/integratedQuery/scoreQuery/thisTermScores/index">&nbsp;&nbsp;
									本学期成绩
								</a>

								<b class="arrow"></b>
							</li>
							
						</ul>

						
					</li>
					
					<li class="" id="1377622">
						<a href="#" class="dropdown-toggle">
							<i class="menu-icon fa fa-caret-right"></i>
							培养方案完成情况
							<b class="arrow fa fa-angle-down"></b>
						</a>

						<b class="arrow"></b>

						
						<ul class="submenu">
							
							<li class="" id="1377623" onclick="toSelect(this);">
								<a href="/student/integratedQuery/planCompletion/index">&nbsp;&nbsp;
									方案完成情况
								</a>

								<b class="arrow"></b>
							</li>
							
						</ul>

						
					</li>
					
					<li class="" id="1443376">
						<a href="#" class="dropdown-toggle">
							<i class="menu-icon fa fa-caret-right"></i>
							指导性计划完成情况
							<b class="arrow fa fa-angle-down"></b>
						</a>

						<b class="arrow"></b>

						
						<ul class="submenu">
							
							<li class="" id="1443377" onclick="toSelect(this);">
								<a href="/student/integratedQuery/instructionPlanQuery/detail/index">&nbsp;&nbsp;
									指导性教学计划
								</a>

								<b class="arrow"></b>
							</li>
							
						</ul>

						
					</li>
					
					<li class="" id="1443378">
						<a href="#" class="dropdown-toggle">
							<i class="menu-icon fa fa-caret-right"></i>
							课程
							<b class="arrow fa fa-angle-down"></b>
						</a>

						<b class="arrow"></b>

						
						<ul class="submenu">
							
							<li class="" id="1443379" onclick="toSelect(this);">
								<a href="/student/integratedQuery/course/courseSchdule/index">&nbsp;&nbsp;
									本学期课程安排
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1443380" onclick="toSelect(this);">
								<a href="/student/integratedQuery/course/courseBasicInformation/basicInf">&nbsp;&nbsp;
									课程基本信息
								</a>

								<b class="arrow"></b>
							</li>
							
						</ul>

						
					</li>
					
					<li class="" id="1443381">
						<a href="#" class="dropdown-toggle">
							<i class="menu-icon fa fa-caret-right"></i>
							教材
							<b class="arrow fa fa-angle-down"></b>
						</a>

						<b class="arrow"></b>

						
						<ul class="submenu">
							
							<li class="" id="1443382" onclick="toSelect(this);">
								<a href="/student/integratedQuery/teachingMaterial/SelectionSearch/index">&nbsp;&nbsp;
									教材选定查询
								</a>

								<b class="arrow"></b>
							</li>
							
							<li class="" id="1443383" onclick="toSelect(this);">
								<a href="/student/integratedQuery/teachingMaterial/ReceiveSearch/index">&nbsp;&nbsp;
									教材领取查询
								</a>

								<b class="arrow"></b>
							</li>
							
						</ul>

						
					</li>
					
				</ul>
				
			</li>
			
			<li class="" id="125803455" onclick="rootMenuClick(this);">
				<a href="#" class="dropdown-toggle">
					<i class="menu-icon fa fa-tag"></i>
					<span class="menu-text">缓考考试
					</span>
					<b class="arrow fa fa-angle-down"></b>
				</a>

				<b class="arrow"></b>
				
				<ul class="submenu" onclick="stopHere();">
					
					<li class="" id="125803456">
						<a href="#" class="dropdown-toggle">
							<i class="menu-icon fa fa-caret-right"></i>
							缓考考试
							<b class="arrow fa fa-angle-down"></b>
						</a>

						<b class="arrow"></b>

						
						<ul class="submenu">
							
							<li class="" id="125803457" onclick="toSelect(this);">
								<a href="/student/delayedExam/control/applyFor/index">&nbsp;&nbsp;
									缓考考试报名
								</a>

								<b class="arrow"></b>
							</li>
							
						</ul>

						
					</li>
					
				</ul>
				
			</li>
			

		</ul>
		<!-- /.nav-list -->

		<div class="sidebar-toggle sidebar-collapse" id="sidebar-collapse">
			<i class="ace-icon fa fa-angle-double-left" data-icon1="ace-icon fa fa-angle-double-left"
			   data-icon2="ace-icon fa fa-angle-double-right"></i>
		</div>

		<script type="text/javascript">
			try {
				ace.settings.check('sidebar', 'collapsed')
			} catch (e) {
			}
		</script>
	</div>

	<div class="main-content">
		<div class="breadcrumbs" id="breadcrumbs">
			<script type="text/javascript">
				try {
					ace.settings.check('breadcrumbs', 'fixed')
				} catch (e) {
				}
			</script>

			<ul class="breadcrumb">
				<li onclick="javascript:window.location.href='/'" style="cursor:pointer;">
					<i class="ace-icon fa fa-home home-icon"></i>
					首页
				</li>
				<li class="active" onclick="ckickTopMenu(this);return false;" id="firmenu">一级菜单</li>
				<li class="active" onclick="ckickTopMenu(this);return false;" id="secmenu">二级菜单</li>
				<li class="active" onclick="ckickTopMenu(this);return false;" id="lastmenu">三级菜单</li>
			</ul>
			<!-- /.breadcrumb -->
		</div>

		<div class="page-content" id="page-content-template">
			<!-- 					<div class="ace-settings-container" id="ace-settings-container"> -->
			<!-- 						<div class="btn btn-app btn-xs btn-warning ace-settings-btn" id="ace-settings-btn"> -->
			<!-- 							<i class="ace-icon fa fa-cog bigger-150"></i> -->
			<!-- 						</div> -->

			<!-- 						<div class="ace-settings-box clearfix" id="ace-settings-box"> -->
			<!-- 							<div class="pull-left width-50"> -->
			<!-- 								<div class="ace-settings-item"> -->
			<!-- 									<div class="pull-left" style="height:13px;"> -->
			<!-- 										<select id="skin-colorpicker" class="hide" style="background: red; height:13px;"> -->
			<!-- 											<option data-skin="no-skin" value="#438EB9">#438EB9</option> -->
			<!-- 											<option data-skin="skin-1" value="#222A2D">#222A2D</option> -->
			<!-- 											<option data-skin="skin-2" value="#C6487E">#C6487E</option> -->
			<!-- 											<option data-skin="skin-3" value="#D0D0D0">#D0D0D0</option> -->
			<!-- 										</select> -->
			<!-- 									</div> -->
			<!-- 									<span>&nbsp; Choose Skin</span> -->
			<!-- 								</div> -->

			<!-- 								<div class="ace-settings-item"> -->
			<!-- 									<input type="checkbox" class="ace ace-checkbox-2" id="ace-settings-navbar" checked="checked"/> -->
			<!-- 									<label class="lbl" for="ace-settings-navbar"> Fixed Navbar</label> -->
			<!-- 								</div> -->

			<!-- 								<div class="ace-settings-item"> -->
			<!-- 									<input type="checkbox" class="ace ace-checkbox-2" id="ace-settings-sidebar" checked="checked"/> -->
			<!-- 									<label class="lbl" for="ace-settings-sidebar"> Fixed Sidebar</label> -->
			<!-- 								</div> -->

			<!-- 								<div class="ace-settings-item"> -->
			<!-- 									<input type="checkbox" class="ace ace-checkbox-2" id="ace-settings-breadcrumbs"/> -->
			<!-- 									<label class="lbl" for="ace-settings-breadcrumbs"> Fixed Breadcrumbs</label> -->
			<!-- 								</div> -->

			<!-- 								<div class="ace-settings-item"> -->
			<!-- 									<input type="checkbox" class="ace ace-checkbox-2" id="ace-settings-rtl" /> -->
			<!-- 									<label class="lbl" for="ace-settings-rtl"> Right To Left (rtl)</label> -->
			<!-- 								</div> -->

			<!-- 								<div class="ace-settings-item"> -->
			<!-- 									<input type="checkbox" class="ace ace-checkbox-2" id="ace-settings-add-container" /> -->
			<!-- 									<label class="lbl" for="ace-settings-add-container"> -->
			<!-- 										Inside -->
			<!-- 										<b>.container</b> -->
			<!-- 									</label> -->
			<!-- 								</div> -->
			<!-- 							</div>/.pull-left -->

			<!-- 							<div class="pull-left width-50"> -->
			<!-- 								<div class="ace-settings-item"> -->
			<!-- 									<input type="checkbox" class="ace ace-checkbox-2" id="ace-settings-hover" /> -->
			<!-- 									<label class="lbl" for="ace-settings-hover"> Submenu on Hover</label> -->
			<!-- 								</div> -->

			<!-- 								<div class="ace-settings-item"> -->
			<!-- 									<input type="checkbox" class="ace ace-checkbox-2" id="ace-settings-compact" /> -->
			<!-- 									<label class="lbl" for="ace-settings-compact"> Compact Sidebar</label> -->
			<!-- 								</div> -->

			<!-- 								<div class="ace-settings-item"> -->
			<!-- 									<input type="checkbox" class="ace ace-checkbox-2" id="ace-settings-highlight" /> -->
			<!-- 									<label class="lbl" for="ace-settings-highlight"> Alt. Active Item</label> -->
			<!-- 								</div> -->
			<!-- 							</div>/.pull-left -->
			<!-- 						</div>/.ace-settings-box -->
			<!-- 					</div>/.ace-settings-container -->

			
<div class=" self-margin">
	
	
	<form id="courseConstruction" action="/student/rollManagement/rollInfo/index" method="GET">
	<!--Main 开始（页面框架） -->
	<div id="left_layout">
		<!-- 查询条件开始2 -->
		<div class="widget">
			<h4 class="header smaller lighter grey">
				<i class="fa fa-cogs"></i>
				与我相关的信息
			</h4>

			<div class="row">
				<div class="col-xs-12 infobox-container" style="text-align:left;">
						
					
						
							<div id="td_div" class="infobox infobox-blue margin-right-10" onclick="pyfa('zx');"
								 style="width: 27%;margin-bottom: 0px;height: 80px!important;">
								<div class="infobox-icon">
									<div class="xc_02"></div>
								</div>
								<div class="infobox-data">
									<span id="td_div" class="infobox-data-number" style="font-size: 15px;">
										
											计算机科学与技术本科培养
											<br> 方案（2015版）
										
										
										<input type="hidden" id="zx" name="zx" value="5223"/>
									</span>

									<div class="infobox-content">主修方案</div>
								</div>
							</div>
						
					

					
						
					
					
						
					
					
					
				</div>
			</div>

			<div class="widget-content">
				
					
				
			</div>

			<h4 class="header smaller lighter grey">
				<i class="glyphicon glyphicon-user"></i>
				基本信息
				<span class="right_top_oper">
					<button id="loading-btn" class="btn btn-purple btn-xs btn-round"
							data-loading-text="正在跳转..." onclick="xiugai();return false;">
						<i class="ace-icon fa fa-pencil-square-o bigger-120"></i>修改信息
					</button>
				</span>
			</h4>
				
			

				
			
				<div>
					<div class="col-xs-4">
						<div class="row">
							<span class="profile-picture" style="padding: 0;">
								<img id="avatar" class="editable img-responsive editable-click editable-empty"
									 style="height:118px;width:97px;"
									 alt="张贤德的照片" title="张贤德的照片" src="/student/rollInfo/img"
									 onerror="this.src='/img/icon/default_photo.png'"></img>
							</span>
						</div>

						<div class="row">
							<div class="self profile-user-info profile-user-info-striped">
								<div class="profile-info-row">
									<div class="profile-info-name">身高</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">体重</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">电子邮件</div>
									<div class="profile-info-value"
										
										 onclick="showInfo('email', this)" style="cursor:pointer; color:blue;">******
										
										
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">QQ号</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">个人简介</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">一卡通账号</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">个人主页</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">手机</div>
									<div class="profile-info-value"
										
										 onclick="showInfo('dh', this)" style="cursor:pointer; color:blue;">******
										
										
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">宿舍电话</div>
									<div class="profile-info-value"
										
										
											>
										
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">邮编</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">血型</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">宿舍地址</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">家长信息</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">备注</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">备注1</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">备注2</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">备注3</div>
									<div class="profile-info-value">
											
									</div>
								</div>
							</div>
						</div>
					</div>

					<div class="col-xs-8" style="border-left:1px solid #DCEBF7">
						<h4 class="header smaller lighter grey">
							<img src="/assets/images/identity.png" style='position:relative;top:-6px;'></img> 学籍信息
						</h4>

						<div>
							<div class="self profile-user-info profile-user-info-striped">
								<div class="profile-info-row">
									<div class="profile-info-name">学号</div>
									<div class="profile-info-value" style="width:34%">
											221171010606
									</div>
									<div class="profile-info-name">姓名</div>
									<div class="profile-info-value">
											张贤德
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">姓名拼音</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">英文姓名</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">证件类型</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">证件号码</div>
									<div class="profile-info-value" id="sfzhDiv"
										
										 onclick="showInfo('sfzh', this)" style="cursor:pointer; color:blue;">******
										
										
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">年级</div>
									<div class="profile-info-value">
											2022级
									</div>
									<div class="profile-info-name">院系</div>
									<div class="profile-info-value">
											计算机与信息工程学院
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">专业</div>
									<div class="profile-info-value">
											计算机科学与技术
									</div>
									<div class="profile-info-name">专业方向</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">班级</div>
									<div class="profile-info-value">
											2022级计算机科学与技术6班
									</div>
									<div class="profile-info-name">校区</div>
									<div class="profile-info-value">
											济南校区
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">辅修专业</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">第二学位专业</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">是否有学籍</div>
									<div class="profile-info-value">
											是
									</div>
									<div class="profile-info-name">是否有国家学籍</div>
									<div class="profile-info-value">
											是
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">学生类别</div>
									<div class="profile-info-value">
											本科
									</div>
									<div class="profile-info-name">学籍状态</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">学科门类</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">特殊学生类型</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">收费类别</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">分流方向</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">培养方式</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">入学日期</div>
									<div class="profile-info-value"
										
										
											>
										
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">因材施教</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">培养层次</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">异动否</div>
									<div class="profile-info-value">
											否
									</div>
									<div class="profile-info-name">是否离校</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">是否应届毕业</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">入学年级</div>
									<div class="profile-info-value">
											2022级
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">学制类型</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">学生证号</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">学生类型</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">是否留学生</div>
									<div class="profile-info-value">
											
									</div>
								</div>
							</div>
						</div>
						<h4 class="header smaller lighter grey">
							<img src="/assets/images/recruit.png" style='position:relative;top:-6px;'></img> 招生信息
						</h4>

						<div>
							<div class="self profile-user-info profile-user-info-striped">
								<div class="profile-info-row">
									<div class="profile-info-name">性别</div>
									<div class="profile-info-value" style="width:34%">
											男
									</div>
									<div class="profile-info-name">民族</div>
									<div class="profile-info-value">
											汉族
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">政治面貌</div>
									<div class="profile-info-value">
											共青团员
									</div>
									<div class="profile-info-name">国籍</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">授课语种</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">出生日期</div>
									<div class="profile-info-value">
											20040711
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">籍贯</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">外语语种</div>
									<div class="profile-info-value">
											英语
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">乘车区间</div>
									<div class="profile-info-value">
										
										
									</div>
									<div class="profile-info-name">考生特征</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">定向委培单位</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">通讯地址</div>
									<div class="profile-info-value">
											山东省济南市历城区荷花路街道办事处坝子833号
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">考区</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">高考考生号</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">高考总分</div>
									<div class="profile-info-value">
											575
									</div>
									<div class="profile-info-name">毕业中学</div>
									<div class="profile-info-value">
											其他
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">入学考试语种</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">录取号</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">录取年份</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">学习形式</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">录取类型</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name" style="background-color: white;"></div>
									<div class="profile-info-value"></div>
								</div>
							</div>
						</div>
						<h4 class="header smaller lighter grey">
							<img src="/assets/images/school.png" style='position:relative;top:-6px;'></img> 毕业信息
						</h4>

						<div>
							<div class="self profile-user-info profile-user-info-striped">
								<div class="profile-info-row">
									<div class="profile-info-name">学位</div>
									<div class="profile-info-value" style="width:34%">
											
									</div>
									<div class="profile-info-name">学位证书编号</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">毕业类型</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">毕业日期</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">预设毕业日期</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">毕业证书编号</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">离校日期</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">授予学位时间</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">电子受理号</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">第二学位</div>
									<div class="profile-info-value">
											
									</div>
								</div>
								<div class="profile-info-row">
									<div class="profile-info-name">第二学位证书编号</div>
									<div class="profile-info-value">
											
									</div>
									<div class="profile-info-name">辅修证书编号</div>
									<div class="profile-info-value">
											
									</div>
								</div>
							</div>
						</div>
					</div>
					<div class='col-xs-12' style="height:20px;"></div>
					<!-- 表单结束 -->
				</div>
			


			<div class='col-xs-12' style="height:20px;"></div>
			<!-- 表单结束 -->
		</div>
	</div>
</div>

<div id="curriculumInfo-divcon">
	<div class="div-title">
		<h3 id="title">奖惩信息</h3>
	</div>
	<div class="row" style="margin-right: 0px; margin-left: 0px;">
		<div class="col-xs-12">
			<div class="row">
				<div class="col-xs-12">
					<table class="table table-striped table-bordered">
						<thead>
						<tr>
							<th>学号</th>
							<th>奖惩日期</th>
							<th>奖惩类别</th>
							<th>奖惩原因</th>
							<th>奖或惩</th>
							<th>备注</th>
						</tr>
						</thead>
						<tbody>
						
						</tbody>
					</table>
				</div>
			</div>
		</div>
	</div>
</div>


<div id="curriculumInfo-divcon1">
	<div class="div-title">
		<h3>学籍异动信息</h3>
	</div>
	<div class="row" style="margin-right: 0px; margin-left: 0px;">
		<div class="col-xs-12">
			<div class="row">
				<div class="col-xs-12">
					<table class="table table-striped table-bordered">
						<thead>
						<tr>
							<th>学号</th>
							<th>异动日期</th>
							<th>异动文号</th>
							<th>异动类别</th>
							<th>异动原因</th>
							<th>异动原因描述</th>
							<th>复学日期</th>
							<th>原院系</th>
							<th>原专业方向</th>
							<th>原班级</th>
							<th>原专业</th>
							<th>原方案计划</th>
						</tr>
						</thead>
						<tbody>
						
						</tbody>
					</table>
				</div>
			</div>
		</div>
	</div>
</div>

<!-- 培养方案 -->
<div id="curriculumInfo-divcon2">
	<div class="div-title">
		<h3>培养方案</h3>
	</div>
	<div class="modal-body no-padding">
		<div class="col-xs-12">
				
			<div class="row" style="margin:0 0;">
				<p><a id="expandAllBtn" style="cursor:pointer;">展开</a> | <a id="collapseAllBtn" style="cursor:pointer;">折叠</a><br></p>

				<div class="col-xs-6" style="border: 1px solid #DDD;padding-left: 0px;padding-right: 0px;">
					<div class="widget-box" style="margin-bottom: 0px;">
						<div class="widget-header">
							<h4 class="widget-title smaller grey" id="tit"></h4>

						</div>
						<div class="widget-body">
							<ul id="treeDemo" class="ztree"></ul>
						</div>
					</div>
				</div>
				<div class="col-xs-6" id="fajh">


				</div>
				<div class="col-xs-6" id="xnxq" style="display: none;">
					<div class="widget-box transparent">
						<div class="widget-header widget-header-small">
							<h4 class="widget-title smaller grey">
								学年学期信息
							</h4>
						</div>
					</div>
					<div class="self profile-user-info profile-user-info-striped">
						<div class="profile-info-row">
							<div class="profile-info-name">计划学年</div>
							<div class="profile-info-value">
								<span class="editable" id="xnxq_jhxn"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">学期类型</div>
							<div class="profile-info-value">
								<span class="editable" id="xnxq_xqlx"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">计划学期</div>
							<div class="profile-info-value">
								<span class="editable" id="xnxq_jhxq"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">最少修读门数</div>
							<div class="profile-info-value">
								<span class="editable" id="xnxq_zsxdms"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">最多修读门数</div>
							<div class="profile-info-value">
								<span class="editable" id="xnxq_zdxdms"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">最少修读学分</div>
							<div class="profile-info-value">
								<span class="editable" id="xnxq_zsxdxf"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">最多修读学分</div>
							<div class="profile-info-value">
								<span class="editable" id="xnxq_zdxdxf"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">最少修读学时</div>
							<div class="profile-info-value">
								<span class="editable" id="xnxq_zsxdxs"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">最多修读学时</div>
							<div class="profile-info-value">
								<span class="editable" id="xnxq_zdxdxs"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">审查要求</div>
							<div class="profile-info-value">
								<span class="editable" id="xnxq_scyq"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">备注</div>
							<div class="profile-info-value">
								<span class="editable" id="xnxq_bz"></span>
							</div>
						</div>
					</div>
				</div>
				<div class="col-xs-6" id="kz" style="display: none;">
					<div class="widget-box transparent">
						<div class="widget-header widget-header-small">
							<h4 class="widget-title smaller grey">
								课组信息
							</h4>
						</div>
					</div>
					<div class="self profile-user-info profile-user-info-striped">
						<div class="profile-info-row">
							<div class="profile-info-name">课组类型</div>
							<div class="profile-info-value">
								<span class="editable" id="kz_kzlx"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">课组名</div>
							<div class="profile-info-value">
								<span class="editable" id="kz_kzm"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">课程总学分</div>
							<div class="profile-info-value">
								<span class="editable" id="kz_kczxf"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">课程总学时</div>
							<div class="profile-info-value">
								<span class="editable" id="kz_kczxs"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">最少修读门数</div>
							<div class="profile-info-value">
								<span class="editable" id="kz_zsms"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">最少修读学分</div>
							<div class="profile-info-value">
								<span class="editable" id="kz_zsxf"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">最少修读学时</div>
							<div class="profile-info-value">
								<span class="editable" id="kz_zsxs"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">审查要求</div>
							<div class="profile-info-value">
								<span class="editable" id="kz_scyq"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">备注</div>
							<div class="profile-info-value">
								<span class="editable" id="kz_bz"></span>
							</div>
						</div>
					</div>
				</div>
				<div class="col-xs-6" id="kc" style="display: none;">
					<div class="widget-box transparent">
						<div class="widget-header widget-header-small">
							<h4 class="widget-title smaller grey">
								课程基本信息
							</h4>
						</div>
					</div>
					<div class="self profile-user-info profile-user-info-striped">
						<div class="profile-info-row">
							<div class="profile-info-name">课程号</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_kch"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">课程名</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_kcm"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">英文课程名</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_ywkcm"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">开课院系</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_xsm"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">开课学期</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_kkxq"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">本研标志</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_bybz"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">学分</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_xf"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">学时</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_xs"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">开始设立日期</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_ksrq"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">结束日期</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_jsrq"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">课程状态</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_kcztmc"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">课内周学时</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_knzxs"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">授课总学时</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_jkzxs"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">设计总学时</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_sjzxs"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">实验总学时</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_syzxs"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">其中上机总学时</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_qzsjzxs"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">讨论辅导总学时</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_tlfdzxs"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">设计作业总学时</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_sjzyzxs"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">课外总学时</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_kwzxs"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">课外学分</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_kwxf"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">课程类别</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_kclbmc"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">教学方式</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_jxfssm"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">教学方式</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_jxfssm"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">教材</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_jc"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">参考书</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_cks"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">师资队伍</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_szdw"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">课程说明</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_kcsm"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">考试类型</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_kslxmc"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">校区</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_xaqm"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">内容简介</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_nrjj"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">备注</div>
							<div class="profile-info-value">
								<span class="editable" id="kc_bz"></span>
							</div>
						</div>
					</div>
				</div>
				<div class="col-xs-6" id="kcfa" style="display: none;">
					<div class="widget-box transparent">
						<div class="widget-header widget-header-small">
							<h4 class="widget-title smaller grey">
								课程方案信息
							</h4>
						</div>
					</div>
					<div class="self profile-user-info profile-user-info-striped">
						<div class="profile-info-row">
							<div class="profile-info-name">计划学年</div>
							<div class="profile-info-value">
								<span class="editable" id="kcfa_xnmc"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">计划学期</div>
							<div class="profile-info-value">
								<span class="editable" id="kcfa_xqm"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">课程属性</div>
							<div class="profile-info-value">
								<span class="editable" id="kcfa_kcsxmc"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">考试类型</div>
							<div class="profile-info-value">
								<span class="editable" id="kcfa_kslxmc"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">备注1</div>
							<div class="profile-info-value">
								<span class="editable" id="kcfa_bz"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">备注2</div>
							<div class="profile-info-value">
								<span class="editable" id="kcfa_bz1"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">备注3</div>
							<div class="profile-info-value">
								<span class="editable" id="kcfa_bz2"></span>
							</div>
						</div>
						<div class="profile-info-row">
							<div class="profile-info-name">备注4</div>
							<div class="profile-info-value">
								<span class="editable" id="kcfa_bz3"></span>
							</div>
						</div>
					</div>
				</div>
			</div>
				
		</div>
	</div>
</div>

</form>


</div>


		</div>
		<!-- /.page-content -->
	</div>
	<!-- /.main-content -->


	<a href="#" id="btn-scroll-up" class="btn-scroll-up btn btn-sm btn-inverse">
		<i class="ace-icon fa fa-angle-double-up icon-only bigger-110"></i>
	</a>
</div>
<!-- /.main-container -->

<!-- inline scripts related to this page -->
<script type="text/javascript">
	$(function () {

		
		if (false) {
			/*var showResetpass = $.cookie('showResetpass');
			if (showResetpass == null) {
				$.cookie('showResetpass', "0");
				var bar = $.cookie('selectionBar');*/
			$("#view-table .modal-dialog").css("width", "60%");
			$('#view-table').modal({
				remote: "/template/forcereset",
				backdrop: "static",
				keyboard: false
			}).on("hidden.bs.modal", function () {
				$("#view-table .modal-dialog").css("width", "60%");
				$("#view-table .modal-content").html('<div class="center">\
                            <img src="/img/icon/pageloading.gif" style="width:28px;height:28px;">\
                        </div>');
				$(this).removeData("bs.modal");
			});


			//}
		}
	});

</script>
<script type="text/javascript">
	$(document).ready(function () {
		var bar = $.cookie('selectionBar');
		if (bar && bar != 0) {
			if (true) {
				var barObj = document.getElementById(bar);
				var pts = $(barObj).parents("li");
				for (var i = pts.length - 1; i >= 0; i--) {
					if (!$(pts[i]).hasClass("open")) {
						$(pts[i]).children("a").click();
					}
				}
				barObj == null ? barObj = document.getElementById(bar) : barObj.className = "active";
				$(barObj).find("a").prepend("<i class='menu-icon fa fa-caret-right'></i>");
			}

			var lastmenu = $("#" + bar).find("a").first().text();
			var secmenu = $("#" + bar).parent().parent().find("a").first().text();
			var firmenu = $("#" + bar).parent().parent().parent().parent().find("a").first().text();

			$("#lastmenu").html(lastmenu.replace(/(^\s*)|(\s*$)/g, ""));
			$("#lastmenu").attr("menuid", bar);
			$("#secmenu").html(secmenu.replace(/(^\s*)|(\s*$)/g, ""));
			$("#secmenu").attr("menuid", $("#" + bar).parent().parent().attr("id"));
			$("#firmenu").html(firmenu.replace(/(^\s*)|(\s*$)/g, ""));
			$("#firmenu").attr("menuid", $("#" + bar).parent().parent().parent().parent().attr("id"));
		} else {
			$("#lastmenu").remove();
			$("#secmenu").remove();
			$("#firmenu").remove();
		}
	});

	function ckickTopMenu(obj) {
		if ($(obj).attr("menuid")) {
			var menuid = $(obj).attr("menuid");
			if ($("#" + menuid + " a:eq(0)").hasClass("dropdown-toggle")) {
				if ($("#" + menuid).parents('.nav-list li').size() > 0) {
					if ($("#" + menuid).parents('.nav-list li:eq(0)').hasClass("open")) {
						$("#" + menuid + " a").trigger("click");
					} else {
						$("#" + menuid).parents('.nav-list li').find("a:eq(0)").trigger("click");
						if (!$("#" + menuid).hasClass("open")) {
							$("#" + menuid + " a").trigger("click");
						}
					}
				} else {
					$("#" + menuid + " a:eq(0)").trigger("click");
				}
			} else {
				window.location.href = $("#" + menuid + " a:eq(0)").attr("href");
			}
		}
	}


	function toSelect(obj) {
		$.cookie('selectionBar', obj.id, {
			path: '/'
		});
	}

	function removeCookie() {
		$.cookie('JSESSIONID', null);
	}

	// 			$(function() {
	// 				$('select').chosen({allow_single_deselect:true});
	// 			});

	function addslidersModel(id, width) {
		var modal = '<div id="' + id + '" class="modal right fade" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true" tabindex="-1">\
	            <div class="modal-dialog">\
	                <div class="modal-content">\
	                    <div class="center">\
                            <img src="/img/icon/pageloading.gif" style="width:28px;height:28px;">\
                        </div>\
	                </div>\
	            </div>\
	        </div>';
		var modal = $(modal).appendTo('body');
		$(".modal-dialog").css("width", width);
		return modal;
	}

	function openWorkRestSchedule() {
		var modal = addslidersModel("work_rest_schedule_modal", "60%");
		var url = "/main/queryRestSchedule";
		modal.modal({
			remote: url
		}).on('hide.bs.modal', function () {
			modal.remove();
		});
	}
</script>
</body>

</html>`

	nameRegex := regexp.MustCompile(`<div class="profile-info-value">\s*([^<]*)\s*</div>`)
	nameMatch := nameRegex.FindAllStringSubmatch(data, -1)

	// 输出匹配结果
	if len(nameMatch) > 0 {
		for i, match := range nameMatch {
			fmt.Println("-------------------------------------------" + strconv.Itoa(i))
			for _, v := range match {
				fmt.Println(v)
				fmt.Println("vvvvvvvvvvvvvvvvvvvv")
			}
		}
	} else {
		fmt.Println("No valid emails found.")
	}
}

func TestRegexp2(t *testing.T) {
	// 示例数据，假设这是一个字符串
	data := `map.set("email", "qq3392313023@163.com");
	map.set("dh", "17353878918");
	map.set("brlxdh", "");
	map.set("sfzh", "370112200407117114");
	map.set("rxrq", "");`
	// 创建正则表达式对象，用于匹配电子邮件地址
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)
	// 创建正则表达式对象，用于匹配 dh 和 sfzh 字段的值
	dhRegex := regexp.MustCompile(`"dh",\s*"(\d{11})"`)
	sfzhRegex := regexp.MustCompile(`"sfzh",\s*"(\d{18})"`)

	// 在字符串中查找所有匹配项
	emailMatch := emailRegex.FindStringSubmatch(data)
	dhMatch := dhRegex.FindStringSubmatch(data)
	sfzhMatch := sfzhRegex.FindStringSubmatch(data)

	// 输出匹配结果
	fmt.Println("Extracted Email:", emailMatch[0])
	fmt.Println("Extracted DH:", dhMatch[1])
	fmt.Println("Extracted SFZH:", sfzhMatch[1])
}
