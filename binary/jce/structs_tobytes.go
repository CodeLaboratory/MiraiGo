// Code generated by internal/generator/jce_gen; DO NOT EDIT.

package jce

func (pkt *RequestPacket) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt16(pkt.IVersion, 1)
	w.WriteByte(pkt.CPacketType, 2)
	w.WriteInt32(pkt.IMessageType, 3)
	w.WriteInt32(pkt.IRequestId, 4)
	w.WriteString(pkt.SServantName, 5)
	w.WriteString(pkt.SFuncName, 6)
	w.WriteBytes(pkt.SBuffer, 7)
	w.WriteInt32(pkt.ITimeout, 8)
	w.writeMapStrStr(pkt.Context, 9)
	w.writeMapStrStr(pkt.Status, 10)
	return w.Bytes()
}

func (pkt *RequestDataVersion3) ToBytes() []byte {
	w := NewJceWriter()
	w.writeMapStrBytes(pkt.Map, 0)
	return w.Bytes()
}

func (pkt *RequestDataVersion2) ToBytes() []byte {
	w := NewJceWriter()
	w.writeMapStrMapStrBytes(pkt.Map, 0)
	return w.Bytes()
}

func (pkt *SsoServerInfo) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteString(pkt.Server, 1)
	w.WriteInt32(pkt.Port, 2)
	w.WriteString(pkt.Location, 8)
	return w.Bytes()
}

func (pkt *FileStoragePushFSSvcList) ToBytes() []byte {
	w := NewJceWriter()
	{ // write pkt.UploadList tag=0
		w.writeHead(9, 0)
		if len(pkt.UploadList) == 0 {
			w.writeHead(12, 0) // w.WriteInt32(0, 0)
		} else {
			w.WriteInt32(int32(len(pkt.UploadList)), 0)
			for _, i := range pkt.UploadList {
				w.writeHead(10, 0)
				w.buf.Write(i.ToBytes())
				w.writeHead(11, 0)
			}
		}
	}
	{ // write pkt.PicDownloadList tag=1
		w.writeHead(9, 1)
		if len(pkt.PicDownloadList) == 0 {
			w.writeHead(12, 0) // w.WriteInt32(0, 0)
		} else {
			w.WriteInt32(int32(len(pkt.PicDownloadList)), 0)
			for _, i := range pkt.PicDownloadList {
				w.writeHead(10, 0)
				w.buf.Write(i.ToBytes())
				w.writeHead(11, 0)
			}
		}
	}
	{ // write pkt.GPicDownloadList tag=2
		w.writeHead(9, 2)
		if len(pkt.GPicDownloadList) == 0 {
			w.writeHead(12, 0) // w.WriteInt32(0, 0)
		} else {
			w.WriteInt32(int32(len(pkt.GPicDownloadList)), 0)
			for _, i := range pkt.GPicDownloadList {
				w.writeHead(10, 0)
				w.buf.Write(i.ToBytes())
				w.writeHead(11, 0)
			}
		}
	}
	{ // write pkt.QZoneProxyServiceList tag=3
		w.writeHead(9, 3)
		if len(pkt.QZoneProxyServiceList) == 0 {
			w.writeHead(12, 0) // w.WriteInt32(0, 0)
		} else {
			w.WriteInt32(int32(len(pkt.QZoneProxyServiceList)), 0)
			for _, i := range pkt.QZoneProxyServiceList {
				w.writeHead(10, 0)
				w.buf.Write(i.ToBytes())
				w.writeHead(11, 0)
			}
		}
	}
	{ // write pkt.UrlEncodeServiceList tag=4
		w.writeHead(9, 4)
		if len(pkt.UrlEncodeServiceList) == 0 {
			w.writeHead(12, 0) // w.WriteInt32(0, 0)
		} else {
			w.WriteInt32(int32(len(pkt.UrlEncodeServiceList)), 0)
			for _, i := range pkt.UrlEncodeServiceList {
				w.writeHead(10, 0)
				w.buf.Write(i.ToBytes())
				w.writeHead(11, 0)
			}
		}
	}
	{ // write BigDataChannel tag=5
		w.writeHead(10, 5)
		w.buf.Write(pkt.BigDataChannel.ToBytes())
		w.writeHead(11, 0)
	}
	{ // write pkt.VipEmotionList tag=6
		w.writeHead(9, 6)
		if len(pkt.VipEmotionList) == 0 {
			w.writeHead(12, 0) // w.WriteInt32(0, 0)
		} else {
			w.WriteInt32(int32(len(pkt.VipEmotionList)), 0)
			for _, i := range pkt.VipEmotionList {
				w.writeHead(10, 0)
				w.buf.Write(i.ToBytes())
				w.writeHead(11, 0)
			}
		}
	}
	{ // write pkt.C2CPicDownList tag=7
		w.writeHead(9, 7)
		if len(pkt.C2CPicDownList) == 0 {
			w.writeHead(12, 0) // w.WriteInt32(0, 0)
		} else {
			w.WriteInt32(int32(len(pkt.C2CPicDownList)), 0)
			for _, i := range pkt.C2CPicDownList {
				w.writeHead(10, 0)
				w.buf.Write(i.ToBytes())
				w.writeHead(11, 0)
			}
		}
	}
	w.WriteBytes(pkt.PttList, 10)
	return w.Bytes()
}

func (pkt *FileStorageServerInfo) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteString(pkt.Server, 1)
	w.WriteInt32(pkt.Port, 2)
	return w.Bytes()
}

func (pkt *BigDataChannel) ToBytes() []byte {
	w := NewJceWriter()
	{ // write pkt.IPLists tag=0
		w.writeHead(9, 0)
		if len(pkt.IPLists) == 0 {
			w.writeHead(12, 0) // w.WriteInt32(0, 0)
		} else {
			w.WriteInt32(int32(len(pkt.IPLists)), 0)
			for _, i := range pkt.IPLists {
				w.writeHead(10, 0)
				w.buf.Write(i.ToBytes())
				w.writeHead(11, 0)
			}
		}
	}
	w.WriteBytes(pkt.SigSession, 1)
	w.WriteBytes(pkt.KeySession, 2)
	w.WriteInt64(pkt.SigUin, 3)
	w.WriteInt32(pkt.ConnectFlag, 4)
	w.WriteBytes(pkt.PbBuf, 5)
	return w.Bytes()
}

func (pkt *BigDataIPList) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.ServiceType, 0)
	{ // write pkt.IPList tag=1
		w.writeHead(9, 1)
		if len(pkt.IPList) == 0 {
			w.writeHead(12, 0) // w.WriteInt32(0, 0)
		} else {
			w.WriteInt32(int32(len(pkt.IPList)), 0)
			for _, i := range pkt.IPList {
				w.writeHead(10, 0)
				w.buf.Write(i.ToBytes())
				w.writeHead(11, 0)
			}
		}
	}
	w.WriteInt64(pkt.FragmentSize, 3)
	return w.Bytes()
}

func (pkt *BigDataIPInfo) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.Type, 0)
	w.WriteString(pkt.Server, 1)
	w.WriteInt64(pkt.Port, 2)
	return w.Bytes()
}

func (pkt *SvcReqRegister) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.Uin, 0)
	w.WriteInt64(pkt.Bid, 1)
	w.WriteByte(pkt.ConnType, 2)
	w.WriteString(pkt.Other, 3)
	w.WriteInt32(pkt.Status, 4)
	w.WriteByte(pkt.OnlinePush, 5)
	w.WriteByte(pkt.IsOnline, 6)
	w.WriteByte(pkt.IsShowOnline, 7)
	w.WriteByte(pkt.KickPC, 8)
	w.WriteByte(pkt.KickWeak, 9)
	w.WriteInt64(pkt.Timestamp, 10)
	w.WriteInt64(pkt.IOSVersion, 11)
	w.WriteByte(pkt.NetType, 12)
	w.WriteString(pkt.BuildVer, 13)
	w.WriteByte(pkt.RegType, 14)
	w.WriteBytes(pkt.DevParam, 15)
	w.WriteBytes(pkt.Guid, 16)
	w.WriteInt32(pkt.LocaleId, 17)
	w.WriteByte(pkt.SilentPush, 18)
	w.WriteString(pkt.DevName, 19)
	w.WriteString(pkt.DevType, 20)
	w.WriteString(pkt.OSVer, 21)
	w.WriteByte(pkt.OpenPush, 22)
	w.WriteInt64(pkt.LargeSeq, 23)
	w.WriteInt64(pkt.LastWatchStartTime, 24)
	w.WriteInt64(pkt.OldSSOIp, 26)
	w.WriteInt64(pkt.NewSSOIp, 27)
	w.WriteString(pkt.ChannelNo, 28)
	w.WriteInt64(pkt.CPID, 29)
	w.WriteString(pkt.VendorName, 30)
	w.WriteString(pkt.VendorOSName, 31)
	w.WriteString(pkt.IOSIdfa, 32)
	w.WriteBytes(pkt.B769, 33)
	w.WriteByte(pkt.IsSetStatus, 34)
	w.WriteBytes(pkt.ServerBuf, 35)
	w.WriteByte(pkt.SetMute, 36)
	w.WriteInt64(pkt.ExtOnlineStatus, 38)
	w.WriteInt32(pkt.BatteryStatus, 39)
	return w.Bytes()
}

func (pkt *SvcRespRegister) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.Uin, 0)
	w.WriteInt64(pkt.Bid, 1)
	w.WriteByte(pkt.ReplyCode, 2)
	w.WriteString(pkt.Result, 3)
	w.WriteInt64(pkt.ServerTime, 4)
	w.WriteByte(pkt.LogQQ, 5)
	w.WriteByte(pkt.NeedKik, 6)
	w.WriteByte(pkt.UpdateFlag, 7)
	w.WriteInt64(pkt.Timestamp, 8)
	w.WriteByte(pkt.CrashFlag, 9)
	w.WriteString(pkt.ClientIp, 10)
	w.WriteInt32(pkt.ClientPort, 11)
	w.WriteInt32(pkt.HelloInterval, 12)
	w.WriteInt32(pkt.LargeSeq, 13)
	w.WriteByte(pkt.LargeSeqUpdate, 14)
	w.WriteBytes(pkt.D769RspBody, 15)
	w.WriteInt32(pkt.Status, 16)
	w.WriteInt64(pkt.ExtOnlineStatus, 17)
	w.WriteInt64(pkt.ClientBatteryGetInterval, 18)
	w.WriteInt64(pkt.ClientAutoStatusInterval, 19)
	return w.Bytes()
}

func (pkt *SvcReqRegisterNew) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.RequestOptional, 0)
	{ // write C2CMsg tag=1
		w.writeHead(10, 1)
		w.buf.Write(pkt.C2CMsg.ToBytes())
		w.writeHead(11, 0)
	}
	{ // write GroupMsg tag=2
		w.writeHead(10, 2)
		w.buf.Write(pkt.GroupMsg.ToBytes())
		w.writeHead(11, 0)
	}
	w.WriteByte(pkt.DisGroupMsgFilter, 14)
	w.WriteByte(pkt.GroupMask, 15)
	w.WriteInt64(pkt.EndSeq, 16)
	w.WriteBytes(pkt.O769Body, 20)
	return w.Bytes()
}

func (pkt *SvcReqGetMsgV2) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.Uin, 0)
	w.WriteInt32(pkt.DateTime, 1)
	w.WriteByte(pkt.RecivePic, 4)
	w.WriteInt16(pkt.Ability, 6)
	w.WriteByte(pkt.Channel, 9)
	w.WriteByte(pkt.Inst, 16)
	w.WriteByte(pkt.ChannelEx, 17)
	w.WriteBytes(pkt.SyncCookie, 18)
	w.WriteInt64(int64(pkt.SyncFlag), 19)
	w.WriteByte(pkt.RambleFlag, 20)
	w.WriteInt64(pkt.GeneralAbi, 26)
	w.WriteBytes(pkt.PubAccountCookie, 27)
	return w.Bytes()
}

func (pkt *SvcReqPullGroupMsgSeq) ToBytes() []byte {
	w := NewJceWriter()
	{ // write pkt.GroupInfo tag=0
		w.writeHead(9, 0)
		if len(pkt.GroupInfo) == 0 {
			w.writeHead(12, 0) // w.WriteInt32(0, 0)
		} else {
			w.WriteInt32(int32(len(pkt.GroupInfo)), 0)
			for _, i := range pkt.GroupInfo {
				w.writeHead(10, 0)
				w.buf.Write(i.ToBytes())
				w.writeHead(11, 0)
			}
		}
	}
	w.WriteByte(pkt.VerifyType, 1)
	w.WriteInt32(pkt.Filter, 2)
	return w.Bytes()
}

func (pkt *PullGroupSeqParam) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.GroupCode, 0)
	w.WriteInt64(pkt.LastSeqId, 1)
	return w.Bytes()
}

func (pkt *SvcRespParam) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt32(pkt.PCStat, 0)
	w.WriteInt32(pkt.IsSupportC2CRoamMsg, 1)
	w.WriteInt32(pkt.IsSupportDataLine, 2)
	w.WriteInt32(pkt.IsSupportPrintable, 3)
	w.WriteInt32(pkt.IsSupportViewPCFile, 4)
	w.WriteInt32(pkt.PcVersion, 5)
	w.WriteInt64(pkt.RoamFlag, 6)
	{ // write pkt.OnlineInfos tag=7
		w.writeHead(9, 7)
		if len(pkt.OnlineInfos) == 0 {
			w.writeHead(12, 0) // w.WriteInt32(0, 0)
		} else {
			w.WriteInt32(int32(len(pkt.OnlineInfos)), 0)
			for _, i := range pkt.OnlineInfos {
				w.writeHead(10, 0)
				w.buf.Write(i.ToBytes())
				w.writeHead(11, 0)
			}
		}
	}
	w.WriteInt32(pkt.PCClientType, 8)
	return w.Bytes()
}

func (pkt *RequestPushNotify) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.Uin, 0)
	w.WriteByte(pkt.Type, 1)
	w.WriteString(pkt.Service, 2)
	w.WriteString(pkt.Cmd, 3)
	w.WriteBytes(pkt.NotifyCookie, 4)
	w.WriteInt32(pkt.MsgType, 5)
	w.WriteInt32(pkt.UserActive, 6)
	w.WriteInt32(pkt.GeneralFlag, 7)
	w.WriteInt64(pkt.BindedUin, 8)
	return w.Bytes()
}

func (pkt *OnlineInfo) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt32(pkt.InstanceId, 0)
	w.WriteInt32(pkt.ClientType, 1)
	w.WriteInt32(pkt.OnlineStatus, 2)
	w.WriteInt32(pkt.PlatformId, 3)
	w.WriteString(pkt.SubPlatform, 4)
	w.WriteInt64(pkt.UClientType, 5)
	return w.Bytes()
}

func (pkt *SvcReqMSFLoginNotify) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.AppId, 0)
	w.WriteByte(pkt.Status, 1)
	w.WriteByte(pkt.Tablet, 2)
	w.WriteInt64(pkt.Platform, 3)
	w.WriteString(pkt.Title, 4)
	w.WriteString(pkt.Info, 5)
	w.WriteInt64(pkt.ProductType, 6)
	w.WriteInt64(pkt.ClientType, 7)
	{ // write pkt.InstanceList tag=8
		w.writeHead(9, 8)
		if len(pkt.InstanceList) == 0 {
			w.writeHead(12, 0) // w.WriteInt32(0, 0)
		} else {
			w.WriteInt32(int32(len(pkt.InstanceList)), 0)
			for _, i := range pkt.InstanceList {
				w.writeHead(10, 0)
				w.buf.Write(i.ToBytes())
				w.writeHead(11, 0)
			}
		}
	}
	return w.Bytes()
}

func (pkt *InstanceInfo) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt32(pkt.AppId, 0)
	w.WriteByte(pkt.Tablet, 1)
	w.WriteInt64(pkt.Platform, 2)
	w.WriteInt64(pkt.ProductType, 3)
	w.WriteInt64(pkt.ClientType, 4)
	return w.Bytes()
}

func (pkt *PushMessageInfo) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.FromUin, 0)
	w.WriteInt64(pkt.MsgTime, 1)
	w.WriteInt16(pkt.MsgType, 2)
	w.WriteInt16(pkt.MsgSeq, 3)
	w.WriteString(pkt.Msg, 4)
	w.WriteInt32(pkt.RealMsgTime, 5)
	w.WriteBytes(pkt.VMsg, 6)
	w.WriteInt64(pkt.AppShareID, 7)
	w.WriteBytes(pkt.MsgCookies, 8)
	w.WriteBytes(pkt.AppShareCookie, 9)
	w.WriteInt64(pkt.MsgUid, 10)
	w.WriteInt64(pkt.LastChangeTime, 11)
	w.WriteInt64(pkt.FromInstId, 14)
	w.WriteBytes(pkt.RemarkOfSender, 15)
	w.WriteString(pkt.FromMobile, 16)
	w.WriteString(pkt.FromName, 17)
	return w.Bytes()
}

func (pkt *SvcRespPushMsg) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.Uin, 0)
	{ // write pkt.DelInfos tag=1
		w.writeHead(9, 1)
		if len(pkt.DelInfos) == 0 {
			w.writeHead(12, 0) // w.WriteInt32(0, 0)
		} else {
			w.WriteInt32(int32(len(pkt.DelInfos)), 0)
			for _, i := range pkt.DelInfos {
				w.writeHead(10, 0)
				w.buf.Write(i.ToBytes())
				w.writeHead(11, 0)
			}
		}
	}
	w.WriteInt32(pkt.Svrip, 2)
	w.WriteBytes(pkt.PushToken, 3)
	w.WriteInt32(pkt.ServiceType, 4)
	return w.Bytes()
}

func (pkt *SvcReqGetDevLoginInfo) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteBytes(pkt.Guid, 0)
	w.WriteString(pkt.AppName, 1)
	w.WriteInt64(pkt.LoginType, 2)
	w.WriteInt64(pkt.Timestamp, 3)
	w.WriteInt64(pkt.NextItemIndex, 4)
	w.WriteInt64(pkt.RequireMax, 5)
	w.WriteInt64(pkt.GetDevListType, 6)
	return w.Bytes()
}

func (pkt *DelMsgInfo) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.FromUin, 0)
	w.WriteInt64(pkt.MsgTime, 1)
	w.WriteInt16(pkt.MsgSeq, 2)
	w.WriteBytes(pkt.MsgCookies, 3)
	w.WriteInt16(pkt.Cmd, 4)
	w.WriteInt64(pkt.MsgType, 5)
	w.WriteInt64(pkt.AppId, 6)
	w.WriteInt64(pkt.SendTime, 7)
	w.WriteInt32(pkt.SsoSeq, 8)
	w.WriteInt32(pkt.SsoIp, 9)
	w.WriteInt32(pkt.ClientIp, 10)
	return w.Bytes()
}

func (pkt *FriendListRequest) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt32(pkt.Reqtype, 0)
	w.WriteByte(pkt.IfReflush, 1)
	w.WriteInt64(pkt.Uin, 2)
	w.WriteInt16(pkt.StartIndex, 3)
	w.WriteInt16(pkt.FriendCount, 4)
	w.WriteByte(pkt.GroupId, 5)
	w.WriteByte(pkt.IfGetGroupInfo, 6)
	w.WriteByte(pkt.GroupStartIndex, 7)
	w.WriteByte(pkt.GroupCount, 8)
	w.WriteByte(pkt.IfGetMSFGroup, 9)
	w.WriteByte(pkt.IfShowTermType, 10)
	w.WriteInt64(pkt.Version, 11)
	w.WriteInt64Slice(pkt.UinList, 12)
	w.WriteInt32(pkt.AppType, 13)
	w.WriteByte(pkt.IfGetDOVId, 14)
	w.WriteByte(pkt.IfGetBothFlag, 15)
	w.WriteBytes(pkt.D50, 16)
	w.WriteBytes(pkt.D6B, 17)
	w.WriteInt64Slice(pkt.SnsTypeList, 18)
	return w.Bytes()
}

func (pkt *FriendInfo) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.FriendUin, 0)
	w.WriteByte(pkt.GroupId, 1)
	w.WriteInt16(pkt.FaceId, 2)
	w.WriteString(pkt.Remark, 3)
	w.WriteByte(pkt.QQType, 4)
	w.WriteByte(pkt.Status, 5)
	w.WriteByte(pkt.MemberLevel, 6)
	w.WriteByte(pkt.IsMqqOnLine, 7)
	w.WriteByte(pkt.QQOnlineState, 8)
	w.WriteByte(pkt.IsIphoneOnline, 9)
	w.WriteByte(pkt.DetailStatusFlag, 10)
	w.WriteByte(pkt.QQOnlineStateV2, 11)
	w.WriteString(pkt.ShowName, 12)
	w.WriteByte(pkt.IsRemark, 13)
	w.WriteString(pkt.Nick, 14)
	w.WriteByte(pkt.SpecialFlag, 15)
	w.WriteBytes(pkt.IMGroupID, 16)
	w.WriteBytes(pkt.MSFGroupID, 17)
	w.WriteInt32(pkt.TermType, 18)
	w.WriteByte(pkt.Network, 20)
	w.WriteBytes(pkt.Ring, 21)
	w.WriteInt64(pkt.AbiFlag, 22)
	w.WriteInt64(pkt.FaceAddonId, 23)
	w.WriteInt32(pkt.NetworkType, 24)
	w.WriteInt64(pkt.VipFont, 25)
	w.WriteInt32(pkt.IconType, 26)
	w.WriteString(pkt.TermDesc, 27)
	w.WriteInt64(pkt.ColorRing, 28)
	w.WriteByte(pkt.ApolloFlag, 29)
	w.WriteInt64(pkt.ApolloTimestamp, 30)
	w.WriteByte(pkt.Sex, 31)
	w.WriteInt64(pkt.FounderFont, 32)
	w.WriteString(pkt.EimId, 33)
	w.WriteString(pkt.EimMobile, 34)
	w.WriteByte(pkt.OlympicTorch, 35)
	w.WriteInt64(pkt.ApolloSignTime, 36)
	w.WriteInt64(pkt.LaviUin, 37)
	w.WriteInt64(pkt.TagUpdateTime, 38)
	w.WriteInt64(pkt.GameLastLoginTime, 39)
	w.WriteInt64(pkt.GameAppId, 40)
	w.WriteBytes(pkt.CardID, 41)
	w.WriteInt64(pkt.BitSet, 42)
	w.WriteByte(pkt.KingOfGloryFlag, 43)
	w.WriteInt64(pkt.KingOfGloryRank, 44)
	w.WriteString(pkt.MasterUin, 45)
	w.WriteInt64(pkt.LastMedalUpdateTime, 46)
	w.WriteInt64(pkt.FaceStoreId, 47)
	w.WriteInt64(pkt.FontEffect, 48)
	w.WriteString(pkt.DOVId, 49)
	w.WriteInt64(pkt.BothFlag, 50)
	w.WriteByte(pkt.CentiShow3DFlag, 51)
	w.WriteBytes(pkt.IntimateInfo, 52)
	w.WriteByte(pkt.ShowNameplate, 53)
	w.WriteByte(pkt.NewLoverDiamondFlag, 54)
	w.WriteBytes(pkt.ExtSnsFrdData, 55)
	w.WriteBytes(pkt.MutualMarkData, 56)
	return w.Bytes()
}

func (pkt *TroopListRequest) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.Uin, 0)
	w.WriteByte(pkt.GetMSFMsgFlag, 1)
	w.WriteBytes(pkt.Cookies, 2)
	w.WriteInt64Slice(pkt.GroupInfo, 3)
	w.WriteByte(pkt.GroupFlagExt, 4)
	w.WriteInt32(pkt.Version, 5)
	w.WriteInt64(pkt.CompanyId, 6)
	w.WriteInt64(pkt.VersionNum, 7)
	w.WriteByte(pkt.GetLongGroupName, 8)
	return w.Bytes()
}

func (pkt *TroopNumber) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.GroupUin, 0)
	w.WriteInt64(pkt.GroupCode, 1)
	w.WriteByte(pkt.Flag, 2)
	w.WriteInt64(pkt.GroupInfoSeq, 3)
	w.WriteString(pkt.GroupName, 4)
	w.WriteString(pkt.GroupMemo, 5)
	w.WriteInt64(pkt.GroupFlagExt, 6)
	w.WriteInt64(pkt.GroupRankSeq, 7)
	w.WriteInt64(pkt.CertificationType, 8)
	w.WriteInt64(pkt.ShutUpTimestamp, 9)
	w.WriteInt64(pkt.MyShutUpTimestamp, 10)
	w.WriteInt64(pkt.CmdUinUinFlag, 11)
	w.WriteInt64(pkt.AdditionalFlag, 12)
	w.WriteInt64(pkt.GroupTypeFlag, 13)
	w.WriteInt64(pkt.GroupSecType, 14)
	w.WriteInt64(pkt.GroupSecTypeInfo, 15)
	w.WriteInt64(pkt.GroupClassExt, 16)
	w.WriteInt64(pkt.AppPrivilegeFlag, 17)
	w.WriteInt64(pkt.SubscriptionUin, 18)
	w.WriteInt64(pkt.MemberNum, 19)
	w.WriteInt64(pkt.MemberNumSeq, 20)
	w.WriteInt64(pkt.MemberCardSeq, 21)
	w.WriteInt64(pkt.GroupFlagExt3, 22)
	w.WriteInt64(pkt.GroupOwnerUin, 23)
	w.WriteByte(pkt.IsConfGroup, 24)
	w.WriteByte(pkt.IsModifyConfGroupFace, 25)
	w.WriteByte(pkt.IsModifyConfGroupName, 26)
	w.WriteInt64(pkt.CmdUinJoinTime, 27)
	w.WriteInt64(pkt.CompanyId, 28)
	w.WriteInt64(pkt.MaxGroupMemberNum, 29)
	w.WriteInt64(pkt.CmdUinGroupMask, 30)
	w.WriteInt64(pkt.GuildAppId, 31)
	w.WriteInt64(pkt.GuildSubType, 32)
	w.WriteInt64(pkt.CmdUinRingtoneID, 33)
	w.WriteInt64(pkt.CmdUinFlagEx2, 34)
	return w.Bytes()
}

func (pkt *TroopMemberListRequest) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.Uin, 0)
	w.WriteInt64(pkt.GroupCode, 1)
	w.WriteInt64(pkt.NextUin, 2)
	w.WriteInt64(pkt.GroupUin, 3)
	w.WriteInt64(pkt.Version, 4)
	w.WriteInt64(pkt.ReqType, 5)
	w.WriteInt64(pkt.GetListAppointTime, 6)
	w.WriteByte(pkt.RichCardNameVer, 7)
	return w.Bytes()
}

func (pkt *TroopMemberInfo) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.MemberUin, 0)
	w.WriteInt16(pkt.FaceId, 1)
	w.WriteByte(pkt.Age, 2)
	w.WriteByte(pkt.Gender, 3)
	w.WriteString(pkt.Nick, 4)
	w.WriteByte(pkt.Status, 5)
	w.WriteString(pkt.ShowName, 6)
	w.WriteString(pkt.Name, 8)
	w.WriteString(pkt.Memo, 12)
	w.WriteString(pkt.AutoRemark, 13)
	w.WriteInt64(pkt.MemberLevel, 14)
	w.WriteInt64(pkt.JoinTime, 15)
	w.WriteInt64(pkt.LastSpeakTime, 16)
	w.WriteInt64(pkt.CreditLevel, 17)
	w.WriteInt64(pkt.Flag, 18)
	w.WriteInt64(pkt.FlagExt, 19)
	w.WriteInt64(pkt.Point, 20)
	w.WriteByte(pkt.Concerned, 21)
	w.WriteByte(pkt.Shielded, 22)
	w.WriteString(pkt.SpecialTitle, 23)
	w.WriteInt64(pkt.SpecialTitleExpireTime, 24)
	w.WriteString(pkt.Job, 25)
	w.WriteByte(pkt.ApolloFlag, 26)
	w.WriteInt64(pkt.ApolloTimestamp, 27)
	w.WriteInt64(pkt.GlobalGroupLevel, 28)
	w.WriteInt64(pkt.TitleId, 29)
	w.WriteInt64(pkt.ShutUpTimestap, 30)
	w.WriteInt64(pkt.GlobalGroupPoint, 31)
	w.WriteByte(pkt.RichCardNameVer, 33)
	w.WriteInt64(pkt.VipType, 34)
	w.WriteInt64(pkt.VipLevel, 35)
	w.WriteInt64(pkt.BigClubLevel, 36)
	w.WriteInt64(pkt.BigClubFlag, 37)
	w.WriteInt64(pkt.Nameplate, 38)
	w.WriteBytes(pkt.GroupHonor, 39)
	return w.Bytes()
}

func (pkt *ModifyGroupCardRequest) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.Zero, 0)
	w.WriteInt64(pkt.GroupCode, 1)
	w.WriteInt64(pkt.NewSeq, 2)
	{ // write pkt.UinInfo tag=3
		w.writeHead(9, 3)
		if len(pkt.UinInfo) == 0 {
			w.writeHead(12, 0) // w.WriteInt32(0, 0)
		} else {
			w.WriteInt32(int32(len(pkt.UinInfo)), 0)
			for _, i := range pkt.UinInfo {
				w.writeHead(10, 0)
				w.buf.Write(i.ToBytes())
				w.writeHead(11, 0)
			}
		}
	}
	return w.Bytes()
}

func (pkt *UinInfo) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.Uin, 0)
	w.WriteInt64(pkt.Flag, 1)
	w.WriteString(pkt.Name, 2)
	w.WriteByte(pkt.Gender, 3)
	w.WriteString(pkt.Phone, 4)
	w.WriteString(pkt.Email, 5)
	w.WriteString(pkt.Remark, 6)
	return w.Bytes()
}

func (pkt *SummaryCardReq) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.Uin, 0)
	w.WriteInt32(pkt.ComeFrom, 1)
	w.WriteInt64(pkt.QzoneFeedTimestamp, 2)
	w.WriteByte(pkt.IsFriend, 3)
	w.WriteInt64(pkt.GroupCode, 4)
	w.WriteInt64(pkt.GroupUin, 5)
	w.WriteInt64(pkt.GetControl, 8)
	w.WriteInt32(pkt.AddFriendSource, 9)
	w.WriteBytes(pkt.SecureSig, 10)
	w.WriteBytesSlice(pkt.ReqServices, 14)
	w.WriteInt64(pkt.TinyId, 15)
	w.WriteInt64(pkt.LikeSource, 16)
	w.WriteByte(pkt.ReqMedalWallInfo, 18)
	w.WriteInt64Slice(pkt.Req0x5ebFieldId, 19)
	w.WriteByte(pkt.ReqNearbyGodInfo, 20)
	w.WriteByte(pkt.ReqExtendCard, 22)
	return w.Bytes()
}

func (pkt *SummaryCardReqSearch) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteString(pkt.Keyword, 0)
	w.WriteString(pkt.CountryCode, 1)
	w.WriteInt32(pkt.Version, 2)
	w.WriteBytesSlice(pkt.ReqServices, 3)
	return w.Bytes()
}

func (pkt *DelFriendReq) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.Uin, 0)
	w.WriteInt64(pkt.DelUin, 1)
	w.WriteByte(pkt.DelType, 2)
	w.WriteInt32(pkt.Version, 3)
	return w.Bytes()
}

func (pkt *VipInfo) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteByte(pkt.Open, 0)
	w.WriteInt32(pkt.Type, 1)
	w.WriteInt32(pkt.Level, 2)
	return w.Bytes()
}

func (pkt *ServiceReqHeader) ToBytes() []byte {
	w := NewJceWriter()
	w.WriteInt64(pkt.Uin, 0)
	w.WriteInt32(pkt.ShVersion, 1)
	w.WriteInt32(pkt.Seq, 2)
	w.WriteInt32(int32(pkt.ReqType), 3)
	w.WriteInt32(int32(pkt.Triggered), 4)
	w.WriteBytes(pkt.Cookies, 5)
	return w.Bytes()
}

func (pkt *ReqFavorite) ToBytes() []byte {
	w := NewJceWriter()
	{
		w.writeHead(10, 0)
		w.buf.Write(pkt.Header.ToBytes())
		w.writeHead(11, 0)
	}
	w.WriteInt64(pkt.Mid, 1)
	w.WriteInt32(pkt.OpType, 2)
	w.WriteInt32(pkt.Source, 3)
	w.WriteInt32(pkt.Times, 4)
	return w.Bytes()
}
