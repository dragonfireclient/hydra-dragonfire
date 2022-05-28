// generated by generate.lua, DO NOT EDIT
package tolua

import (
	"github.com/anon55555/mt"
	"github.com/yuin/gopher-lua"
)

func AnimType(l *lua.LState, val mt.AnimType) lua.LValue {
	switch val {
	case mt.NoAnim:
		return lua.LNil
	case mt.VerticalFrameAnim:
		return lua.LString("vertical_frame")
	case mt.SpriteSheetAnim:
		return lua.LString("sprite_sheet")
	}
	panic("impossible")
	return lua.LNil
}

func ChatMsgType(l *lua.LState, val mt.ChatMsgType) lua.LValue {
	switch val {
	case mt.RawMsg:
		return lua.LString("raw")
	case mt.NormalMsg:
		return lua.LString("normal")
	case mt.AnnounceMsg:
		return lua.LString("announce")
	case mt.SysMsg:
		return lua.LString("sys")
	}
	panic("impossible")
	return lua.LNil
}

func HotbarParam(l *lua.LState, val mt.HotbarParam) lua.LValue {
	switch val {
	case mt.HotbarSize:
		return lua.LString("size")
	case mt.HotbarImg:
		return lua.LString("img")
	case mt.HotbarSelImg:
		return lua.LString("sel_img")
	}
	panic("impossible")
	return lua.LNil
}

func HUDField(l *lua.LState, val mt.HUDField) lua.LValue {
	switch val {
	case mt.HUDPos:
		return lua.LString("pos")
	case mt.HUDName:
		return lua.LString("name")
	case mt.HUDScale:
		return lua.LString("scale")
	case mt.HUDText:
		return lua.LString("text")
	case mt.HUDNumber:
		return lua.LString("number")
	case mt.HUDItem:
		return lua.LString("item")
	case mt.HUDDir:
		return lua.LString("dir")
	case mt.HUDAlign:
		return lua.LString("align")
	case mt.HUDOffset:
		return lua.LString("offset")
	case mt.HUDWorldPos:
		return lua.LString("world_pos")
	case mt.HUDSize:
		return lua.LString("size")
	case mt.HUDZIndex:
		return lua.LString("z_index")
	case mt.HUDText2:
		return lua.LString("text_2")
	}
	panic("impossible")
	return lua.LNil
}

func HUDType(l *lua.LState, val mt.HUDType) lua.LValue {
	switch val {
	case mt.ImgHUD:
		return lua.LString("img")
	case mt.TextHUD:
		return lua.LString("text")
	case mt.StatbarHUD:
		return lua.LString("statbar")
	case mt.InvHUD:
		return lua.LString("inv")
	case mt.WaypointHUD:
		return lua.LString("waypoint")
	case mt.ImgWaypointHUD:
		return lua.LString("img_waypoint")
	}
	panic("impossible")
	return lua.LNil
}

func KickReason(l *lua.LState, val mt.KickReason) lua.LValue {
	switch val {
	case mt.WrongPasswd:
		return lua.LString("wrong_passwd")
	case mt.UnexpectedData:
		return lua.LString("unexpected_data")
	case mt.SrvIsSingleplayer:
		return lua.LString("srv_is_singleplayer")
	case mt.UnsupportedVer:
		return lua.LString("unsupported_ver")
	case mt.BadNameChars:
		return lua.LString("bad_name_chars")
	case mt.BadName:
		return lua.LString("bad_name")
	case mt.TooManyClts:
		return lua.LString("too_many_clts")
	case mt.EmptyPasswd:
		return lua.LString("empty_passwd")
	case mt.AlreadyConnected:
		return lua.LString("already_connected")
	case mt.SrvErr:
		return lua.LString("srv_err")
	case mt.Custom:
		return lua.LString("custom")
	case mt.Shutdown:
		return lua.LString("shutdown")
	case mt.Crash:
		return lua.LString("crash")
	}
	panic("impossible")
	return lua.LNil
}

func ModChanSig(l *lua.LState, val mt.ModChanSig) lua.LValue {
	switch val {
	case mt.JoinOK:
		return lua.LString("join_ok")
	case mt.JoinFail:
		return lua.LString("join_fail")
	case mt.LeaveOK:
		return lua.LString("leave_ok")
	case mt.LeaveFail:
		return lua.LString("leave_fail")
	case mt.NotRegistered:
		return lua.LString("not_registered")
	case mt.SetState:
		return lua.LString("set_state")
	}
	panic("impossible")
	return lua.LNil
}

func PlayerListUpdateType(l *lua.LState, val mt.PlayerListUpdateType) lua.LValue {
	switch val {
	case mt.InitPlayers:
		return lua.LString("init")
	case mt.AddPlayers:
		return lua.LString("add")
	case mt.RemovePlayers:
		return lua.LString("remove")
	}
	panic("impossible")
	return lua.LNil
}

func SoundSrcType(l *lua.LState, val mt.SoundSrcType) lua.LValue {
	switch val {
	case mt.NoSrc:
		return lua.LNil
	case mt.PosSrc:
		return lua.LString("pos")
	case mt.AOSrc:
		return lua.LString("ao")
	}
	panic("impossible")
	return lua.LNil
}

func AuthMethods(l *lua.LState, val mt.AuthMethods) lua.LValue {
	tbl := l.NewTable()
	if val&mt.LegacyPasswd != 0 {
		l.SetField(tbl, "legacy_passwd", lua.LTrue)
	}
	if val&mt.SRP != 0 {
		l.SetField(tbl, "srp", lua.LTrue)
	}
	if val&mt.FirstSRP != 0 {
		l.SetField(tbl, "first_srp", lua.LTrue)
	}
	return tbl
}

func CSMRestrictionFlags(l *lua.LState, val mt.CSMRestrictionFlags) lua.LValue {
	tbl := l.NewTable()
	if val&mt.NoCSMs != 0 {
		l.SetField(tbl, "no_csms", lua.LTrue)
	}
	if val&mt.NoChatMsgs != 0 {
		l.SetField(tbl, "no_chat_msgs", lua.LTrue)
	}
	if val&mt.NoNodeDefs != 0 {
		l.SetField(tbl, "no_node_defs", lua.LTrue)
	}
	if val&mt.LimitMapRange != 0 {
		l.SetField(tbl, "limit_map_range", lua.LTrue)
	}
	if val&mt.NoPlayerList != 0 {
		l.SetField(tbl, "no_player_list", lua.LTrue)
	}
	return tbl
}

func HUDFlags(l *lua.LState, val mt.HUDFlags) lua.LValue {
	tbl := l.NewTable()
	if val&mt.ShowHotbar != 0 {
		l.SetField(tbl, "hotbar", lua.LTrue)
	}
	if val&mt.ShowHealthBar != 0 {
		l.SetField(tbl, "health_bar", lua.LTrue)
	}
	if val&mt.ShowCrosshair != 0 {
		l.SetField(tbl, "crosshair", lua.LTrue)
	}
	if val&mt.ShowWieldedItem != 0 {
		l.SetField(tbl, "wielded_item", lua.LTrue)
	}
	if val&mt.ShowBreathBar != 0 {
		l.SetField(tbl, "breath_bar", lua.LTrue)
	}
	if val&mt.ShowMinimap != 0 {
		l.SetField(tbl, "minimap", lua.LTrue)
	}
	if val&mt.ShowRadarMinimap != 0 {
		l.SetField(tbl, "radar_minimap", lua.LTrue)
	}
	return tbl
}

func HUD(l *lua.LState, val mt.HUD) lua.LValue {
	tbl := l.NewTable()
	l.SetField(tbl, "align", Vec2(l, [2]lua.LNumber{lua.LNumber(val.Align[0]), lua.LNumber(val.Align[1])}))
	l.SetField(tbl, "dir", lua.LNumber(val.Dir))
	l.SetField(tbl, "item", lua.LNumber(val.Item))
	l.SetField(tbl, "name", lua.LString(string(val.Name)))
	l.SetField(tbl, "number", lua.LNumber(val.Number))
	l.SetField(tbl, "offset", Vec2(l, [2]lua.LNumber{lua.LNumber(val.Offset[0]), lua.LNumber(val.Offset[1])}))
	l.SetField(tbl, "pos", Vec2(l, [2]lua.LNumber{lua.LNumber(val.Pos[0]), lua.LNumber(val.Pos[1])}))
	l.SetField(tbl, "scale", Vec2(l, [2]lua.LNumber{lua.LNumber(val.Scale[0]), lua.LNumber(val.Scale[1])}))
	l.SetField(tbl, "size", Vec2(l, [2]lua.LNumber{lua.LNumber(val.Size[0]), lua.LNumber(val.Size[1])}))
	l.SetField(tbl, "text", lua.LString(string(val.Text)))
	l.SetField(tbl, "text_2", lua.LString(string(val.Text2)))
	l.SetField(tbl, "type", HUDType(l, val.Type))
	l.SetField(tbl, "world_pos", Vec3(l, [3]lua.LNumber{lua.LNumber(val.WorldPos[0]), lua.LNumber(val.WorldPos[1]), lua.LNumber(val.WorldPos[2])}))
	l.SetField(tbl, "z_index", lua.LNumber(val.ZIndex))
	return tbl
}

func Node(l *lua.LState, val mt.Node) lua.LValue {
	tbl := l.NewTable()
	l.SetField(tbl, "param0", lua.LNumber(val.Param0))
	l.SetField(tbl, "param1", lua.LNumber(val.Param1))
	l.SetField(tbl, "param2", lua.LNumber(val.Param2))
	return tbl
}

func TileAnim(l *lua.LState, val mt.TileAnim) lua.LValue {
	tbl := l.NewTable()
	l.SetField(tbl, "aspect_ratio", Vec2(l, [2]lua.LNumber{lua.LNumber(val.AspectRatio[0]), lua.LNumber(val.AspectRatio[1])}))
	l.SetField(tbl, "duration", lua.LNumber(val.Duration))
	l.SetField(tbl, "n_frames", Vec2(l, [2]lua.LNumber{lua.LNumber(val.NFrames[0]), lua.LNumber(val.NFrames[1])}))
	l.SetField(tbl, "type", AnimType(l, val.Type))
	return tbl
}

func PktType(pkt *mt.Pkt) lua.LString {
	switch pkt.Cmd.(type) {
	case *mt.ToCltAcceptAuth:
		return lua.LString("accept_auth")
	case *mt.ToCltAcceptSudoMode:
		return lua.LString("accept_sudo_mode")
	case *mt.ToCltAddHUD:
		return lua.LString("add_hud")
	case *mt.ToCltAddNode:
		return lua.LString("add_node")
	case *mt.ToCltAddParticleSpawner:
		return lua.LString("add_particle_spawner")
	case *mt.ToCltAddPlayerVel:
		return lua.LString("add_player_vel")
	case *mt.ToCltAnnounceMedia:
		return lua.LString("announce_media")
	case *mt.ToCltAOMsgs:
		return lua.LString("ao_msgs")
	case *mt.ToCltAORmAdd:
		return lua.LString("ao_rm_add")
	case *mt.ToCltBlkData:
		return lua.LString("blk_data")
	case *mt.ToCltBreath:
		return lua.LString("breath")
	case *mt.ToCltChangeHUD:
		return lua.LString("change_hud")
	case *mt.ToCltChatMsg:
		return lua.LString("chat_msg")
	case *mt.ToCltCloudParams:
		return lua.LString("cloud_params")
	case *mt.ToCltCSMRestrictionFlags:
		return lua.LString("csm_restriction_flags")
	case *mt.ToCltDeathScreen:
		return lua.LString("death_screen")
	case *mt.ToCltDelParticleSpawner:
		return lua.LString("del_particle_spawner")
	case *mt.ToCltDenySudoMode:
		return lua.LString("deny_sudo_mode")
	case *mt.ToCltDetachedInv:
		return lua.LString("detached_inv")
	case *mt.ToCltDisco:
		return lua.LString("disco")
	case *mt.ToCltEyeOffset:
		return lua.LString("eye_offset")
	case *mt.ToCltFadeSound:
		return lua.LString("fade_sound")
	case *mt.ToCltFormspecPrepend:
		return lua.LString("formspec_prepend")
	case *mt.ToCltFOV:
		return lua.LString("fov")
	case *mt.ToCltHello:
		return lua.LString("hello")
	case *mt.ToCltHP:
		return lua.LString("hp")
	case *mt.ToCltHUDFlags:
		return lua.LString("hud_flags")
	case *mt.ToCltInv:
		return lua.LString("inv")
	case *mt.ToCltInvFormspec:
		return lua.LString("inv_formspec")
	case *mt.ToCltItemDefs:
		return lua.LString("item_defs")
	case *mt.ToCltKick:
		return lua.LString("kick")
	case *mt.ToCltLegacyKick:
		return lua.LString("legacy_kick")
	case *mt.ToCltLocalPlayerAnim:
		return lua.LString("local_player_anim")
	case *mt.ToCltMedia:
		return lua.LString("media")
	case *mt.ToCltMediaPush:
		return lua.LString("media_push")
	case *mt.ToCltMinimapModes:
		return lua.LString("minimap_modes")
	case *mt.ToCltModChanMsg:
		return lua.LString("mod_chan_msg")
	case *mt.ToCltModChanSig:
		return lua.LString("mod_chan_sig")
	case *mt.ToCltMoonParams:
		return lua.LString("moon_params")
	case *mt.ToCltMovePlayer:
		return lua.LString("move_player")
	case *mt.ToCltMovement:
		return lua.LString("movement")
	case *mt.ToCltNodeDefs:
		return lua.LString("node_defs")
	case *mt.ToCltNodeMetasChanged:
		return lua.LString("node_metas_changed")
	case *mt.ToCltOverrideDayNightRatio:
		return lua.LString("override_day_night_ratio")
	case *mt.ToCltPlaySound:
		return lua.LString("play_sound")
	case *mt.ToCltPrivs:
		return lua.LString("privs")
	case *mt.ToCltRemoveNode:
		return lua.LString("remove_node")
	case *mt.ToCltRmHUD:
		return lua.LString("rm_hud")
	case *mt.ToCltSetHotbarParam:
		return lua.LString("set_hotbar_param")
	case *mt.ToCltShowFormspec:
		return lua.LString("show_formspec")
	case *mt.ToCltSkyParams:
		return lua.LString("sky_params")
	case *mt.ToCltSpawnParticle:
		return lua.LString("spawn_particle")
	case *mt.ToCltSRPBytesSaltB:
		return lua.LString("srp_bytes_salt_b")
	case *mt.ToCltStarParams:
		return lua.LString("star_params")
	case *mt.ToCltStopSound:
		return lua.LString("stop_sound")
	case *mt.ToCltSunParams:
		return lua.LString("sun_params")
	case *mt.ToCltTimeOfDay:
		return lua.LString("time_of_day")
	case *mt.ToCltUpdatePlayerList:
		return lua.LString("update_player_list")
	}
	panic("impossible")
	return ""
}

func Pkt(l *lua.LState, pkt *mt.Pkt) lua.LValue {
	if pkt == nil {
		return lua.LNil
	}
	tbl := l.NewTable()
	l.SetField(tbl, "_type", PktType(pkt))
	switch val := pkt.Cmd.(type) {
	case *mt.ToCltAcceptAuth:
		l.SetField(tbl, "map_seed", lua.LNumber(val.MapSeed))
		l.SetField(tbl, "player_pos", Vec3(l, [3]lua.LNumber{lua.LNumber(val.PlayerPos[0]), lua.LNumber(val.PlayerPos[1]), lua.LNumber(val.PlayerPos[2])}))
		l.SetField(tbl, "send_interval", lua.LNumber(val.SendInterval))
		l.SetField(tbl, "sudo_auth_methods", AuthMethods(l, val.SudoAuthMethods))
	case *mt.ToCltAddHUD:
		l.SetField(tbl, "hud", HUD(l, val.HUD))
		l.SetField(tbl, "id", lua.LNumber(val.ID))
	case *mt.ToCltAddNode:
		l.SetField(tbl, "keep_meta", lua.LBool(val.KeepMeta))
		l.SetField(tbl, "node", Node(l, val.Node))
		l.SetField(tbl, "pos", Vec3(l, [3]lua.LNumber{lua.LNumber(val.Pos[0]), lua.LNumber(val.Pos[1]), lua.LNumber(val.Pos[2])}))
	case *mt.ToCltAddParticleSpawner:
		l.SetField(tbl, "acc", Box3(l, [2][3]lua.LNumber{{lua.LNumber(val.Acc[0][0]), lua.LNumber(val.Acc[0][1]), lua.LNumber(val.Acc[0][2])}, {lua.LNumber(val.Acc[1][0]), lua.LNumber(val.Acc[1][1]), lua.LNumber(val.Acc[1][2])}}))
		l.SetField(tbl, "amount", lua.LNumber(val.Amount))
		l.SetField(tbl, "anim_params", TileAnim(l, val.AnimParams))
		l.SetField(tbl, "ao_collision", lua.LBool(val.AOCollision))
		l.SetField(tbl, "collide", lua.LBool(val.Collide))
		l.SetField(tbl, "collision_rm", lua.LBool(val.CollisionRm))
		l.SetField(tbl, "duration", lua.LNumber(val.Duration))
		l.SetField(tbl, "expiration_time", Box1(l, [2]lua.LNumber{lua.LNumber(val.ExpirationTime[0]), lua.LNumber(val.ExpirationTime[1])}))
		l.SetField(tbl, "glow", lua.LNumber(val.Glow))
		l.SetField(tbl, "id", lua.LNumber(val.ID))
		l.SetField(tbl, "node_param0", lua.LNumber(val.NodeParam0))
		l.SetField(tbl, "node_param2", lua.LNumber(val.NodeParam2))
		l.SetField(tbl, "node_tile", lua.LNumber(val.NodeTile))
		l.SetField(tbl, "pos", Box3(l, [2][3]lua.LNumber{{lua.LNumber(val.Pos[0][0]), lua.LNumber(val.Pos[0][1]), lua.LNumber(val.Pos[0][2])}, {lua.LNumber(val.Pos[1][0]), lua.LNumber(val.Pos[1][1]), lua.LNumber(val.Pos[1][2])}}))
		l.SetField(tbl, "size", Box1(l, [2]lua.LNumber{lua.LNumber(val.Size[0]), lua.LNumber(val.Size[1])}))
		l.SetField(tbl, "texture", lua.LString(string(val.Texture)))
		l.SetField(tbl, "vel", Box3(l, [2][3]lua.LNumber{{lua.LNumber(val.Vel[0][0]), lua.LNumber(val.Vel[0][1]), lua.LNumber(val.Vel[0][2])}, {lua.LNumber(val.Vel[1][0]), lua.LNumber(val.Vel[1][1]), lua.LNumber(val.Vel[1][2])}}))
		l.SetField(tbl, "vertical", lua.LBool(val.Vertical))
	case *mt.ToCltAddPlayerVel:
		l.SetField(tbl, "vel", Vec3(l, [3]lua.LNumber{lua.LNumber(val.Vel[0]), lua.LNumber(val.Vel[1]), lua.LNumber(val.Vel[2])}))
	case *mt.ToCltBlkData:
		l.SetField(tbl, "blkpos", Vec3(l, [3]lua.LNumber{lua.LNumber(val.Blkpos[0]), lua.LNumber(val.Blkpos[1]), lua.LNumber(val.Blkpos[2])}))
	case *mt.ToCltBreath:
		l.SetField(tbl, "breath", lua.LNumber(val.Breath))
	case *mt.ToCltChangeHUD:
		if val.Field == mt.HUDAlign {
			l.SetField(tbl, "align", Vec2(l, [2]lua.LNumber{lua.LNumber(val.Align[0]), lua.LNumber(val.Align[1])}))
		}
		if val.Field == mt.HUDDir {
			l.SetField(tbl, "dir", lua.LNumber(val.Dir))
		}
		l.SetField(tbl, "field", HUDField(l, val.Field))
		l.SetField(tbl, "id", lua.LNumber(val.ID))
		if val.Field == mt.HUDItem {
			l.SetField(tbl, "item", lua.LNumber(val.Item))
		}
		if val.Field == mt.HUDName {
			l.SetField(tbl, "name", lua.LString(string(val.Name)))
		}
		if val.Field == mt.HUDNumber {
			l.SetField(tbl, "number", lua.LNumber(val.Number))
		}
		if val.Field == mt.HUDOffset {
			l.SetField(tbl, "offset", Vec2(l, [2]lua.LNumber{lua.LNumber(val.Offset[0]), lua.LNumber(val.Offset[1])}))
		}
		if val.Field == mt.HUDPos {
			l.SetField(tbl, "pos", Vec2(l, [2]lua.LNumber{lua.LNumber(val.Pos[0]), lua.LNumber(val.Pos[1])}))
		}
		if val.Field == mt.HUDSize {
			l.SetField(tbl, "size", Vec2(l, [2]lua.LNumber{lua.LNumber(val.Size[0]), lua.LNumber(val.Size[1])}))
		}
		if val.Field == mt.HUDText {
			l.SetField(tbl, "text", lua.LString(string(val.Text)))
		}
		if val.Field == mt.HUDText2 {
			l.SetField(tbl, "text_2", lua.LString(string(val.Text2)))
		}
		if val.Field == mt.HUDWorldPos {
			l.SetField(tbl, "world_pos", Vec3(l, [3]lua.LNumber{lua.LNumber(val.WorldPos[0]), lua.LNumber(val.WorldPos[1]), lua.LNumber(val.WorldPos[2])}))
		}
		if val.Field == mt.HUDZIndex {
			l.SetField(tbl, "z_index", lua.LNumber(val.ZIndex))
		}
	case *mt.ToCltChatMsg:
		l.SetField(tbl, "sender", lua.LString(string(val.Sender)))
		l.SetField(tbl, "text", lua.LString(string(val.Text)))
		l.SetField(tbl, "timestamp", lua.LNumber(val.Timestamp))
		l.SetField(tbl, "type", ChatMsgType(l, val.Type))
	case *mt.ToCltCloudParams:
		l.SetField(tbl, "ambient_color", Color(l, val.AmbientColor))
		l.SetField(tbl, "density", lua.LNumber(val.Density))
		l.SetField(tbl, "diffuse_color", Color(l, val.DiffuseColor))
		l.SetField(tbl, "height", lua.LNumber(val.Height))
		l.SetField(tbl, "speed", Vec2(l, [2]lua.LNumber{lua.LNumber(val.Speed[0]), lua.LNumber(val.Speed[1])}))
		l.SetField(tbl, "thickness", lua.LNumber(val.Thickness))
	case *mt.ToCltCSMRestrictionFlags:
		l.SetField(tbl, "flags", CSMRestrictionFlags(l, val.Flags))
		l.SetField(tbl, "map_range", lua.LNumber(val.MapRange))
	case *mt.ToCltDeathScreen:
		l.SetField(tbl, "point_at", Vec3(l, [3]lua.LNumber{lua.LNumber(val.PointAt[0]), lua.LNumber(val.PointAt[1]), lua.LNumber(val.PointAt[2])}))
		l.SetField(tbl, "point_cam", lua.LBool(val.PointCam))
	case *mt.ToCltDelParticleSpawner:
		l.SetField(tbl, "id", lua.LNumber(val.ID))
	case *mt.ToCltDetachedInv:
		l.SetField(tbl, "inv", lua.LString(string(val.Inv)))
		l.SetField(tbl, "keep", lua.LBool(val.Keep))
		l.SetField(tbl, "len", lua.LNumber(val.Len))
		l.SetField(tbl, "name", lua.LString(string(val.Name)))
	case *mt.ToCltEyeOffset:
		l.SetField(tbl, "first", Vec3(l, [3]lua.LNumber{lua.LNumber(val.First[0]), lua.LNumber(val.First[1]), lua.LNumber(val.First[2])}))
		l.SetField(tbl, "third", Vec3(l, [3]lua.LNumber{lua.LNumber(val.Third[0]), lua.LNumber(val.Third[1]), lua.LNumber(val.Third[2])}))
	case *mt.ToCltFadeSound:
		l.SetField(tbl, "gain", lua.LNumber(val.Gain))
		l.SetField(tbl, "id", lua.LNumber(val.ID))
		l.SetField(tbl, "step", lua.LNumber(val.Step))
	case *mt.ToCltFormspecPrepend:
		l.SetField(tbl, "prepend", lua.LString(string(val.Prepend)))
	case *mt.ToCltFOV:
		l.SetField(tbl, "fov", lua.LNumber(val.FOV))
		l.SetField(tbl, "multiplier", lua.LBool(val.Multiplier))
		l.SetField(tbl, "transition_time", lua.LNumber(val.TransitionTime))
	case *mt.ToCltHello:
		l.SetField(tbl, "auth_methods", AuthMethods(l, val.AuthMethods))
		l.SetField(tbl, "compression", lua.LNumber(val.Compression))
		l.SetField(tbl, "proto_ver", lua.LNumber(val.ProtoVer))
		l.SetField(tbl, "serialize_ver", lua.LNumber(val.SerializeVer))
		l.SetField(tbl, "username", lua.LString(string(val.Username)))
	case *mt.ToCltHP:
		l.SetField(tbl, "hp", lua.LNumber(val.HP))
	case *mt.ToCltHUDFlags:
		l.SetField(tbl, "flags", HUDFlags(l, val.Flags))
		l.SetField(tbl, "mask", HUDFlags(l, val.Mask))
	case *mt.ToCltInv:
		l.SetField(tbl, "inv", lua.LString(string(val.Inv)))
	case *mt.ToCltInvFormspec:
		l.SetField(tbl, "formspec", lua.LString(string(val.Formspec)))
	case *mt.ToCltKick:
		if val.Reason == mt.Custom || val.Reason == mt.Shutdown || val.Reason == mt.Crash {
			l.SetField(tbl, "custom", lua.LString(string(val.Custom)))
		}
		l.SetField(tbl, "reason", KickReason(l, val.Reason))
		if val.Reason == mt.Shutdown || val.Reason == mt.Crash {
			l.SetField(tbl, "reconnect", lua.LBool(val.Reconnect))
		}
	case *mt.ToCltLegacyKick:
		l.SetField(tbl, "reason", lua.LString(string(val.Reason)))
	case *mt.ToCltLocalPlayerAnim:
		l.SetField(tbl, "dig", Box1(l, [2]lua.LNumber{lua.LNumber(val.Dig[0]), lua.LNumber(val.Dig[1])}))
		l.SetField(tbl, "idle", Box1(l, [2]lua.LNumber{lua.LNumber(val.Idle[0]), lua.LNumber(val.Idle[1])}))
		l.SetField(tbl, "speed", lua.LNumber(val.Speed))
		l.SetField(tbl, "walk", Box1(l, [2]lua.LNumber{lua.LNumber(val.Walk[0]), lua.LNumber(val.Walk[1])}))
		l.SetField(tbl, "walk_dig", Box1(l, [2]lua.LNumber{lua.LNumber(val.WalkDig[0]), lua.LNumber(val.WalkDig[1])}))
	case *mt.ToCltMediaPush:
		l.SetField(tbl, "data", lua.LString(string(val.Data)))
		l.SetField(tbl, "filename", lua.LString(string(val.Filename)))
		l.SetField(tbl, "sha1", lua.LString(string(val.SHA1[:])))
		l.SetField(tbl, "should_cache", lua.LBool(val.ShouldCache))
	case *mt.ToCltModChanMsg:
		l.SetField(tbl, "channel", lua.LString(string(val.Channel)))
		l.SetField(tbl, "msg", lua.LString(string(val.Msg)))
		l.SetField(tbl, "sender", lua.LString(string(val.Sender)))
	case *mt.ToCltModChanSig:
		l.SetField(tbl, "channel", lua.LString(string(val.Channel)))
		l.SetField(tbl, "signal", ModChanSig(l, val.Signal))
	case *mt.ToCltMoonParams:
		l.SetField(tbl, "size", lua.LNumber(val.Size))
		l.SetField(tbl, "texture", lua.LString(string(val.Texture)))
		l.SetField(tbl, "tone_map", lua.LString(string(val.ToneMap)))
		l.SetField(tbl, "visible", lua.LBool(val.Visible))
	case *mt.ToCltMovePlayer:
		l.SetField(tbl, "pitch", lua.LNumber(val.Pitch))
		l.SetField(tbl, "pos", Vec3(l, [3]lua.LNumber{lua.LNumber(val.Pos[0]), lua.LNumber(val.Pos[1]), lua.LNumber(val.Pos[2])}))
		l.SetField(tbl, "yaw", lua.LNumber(val.Yaw))
	case *mt.ToCltMovement:
		l.SetField(tbl, "air_accel", lua.LNumber(val.AirAccel))
		l.SetField(tbl, "climb_speed", lua.LNumber(val.ClimbSpeed))
		l.SetField(tbl, "crouch_speed", lua.LNumber(val.CrouchSpeed))
		l.SetField(tbl, "default_accel", lua.LNumber(val.DefaultAccel))
		l.SetField(tbl, "fast_accel", lua.LNumber(val.FastAccel))
		l.SetField(tbl, "fast_speed", lua.LNumber(val.FastSpeed))
		l.SetField(tbl, "fluidity", lua.LNumber(val.Fluidity))
		l.SetField(tbl, "gravity", lua.LNumber(val.Gravity))
		l.SetField(tbl, "jump_speed", lua.LNumber(val.JumpSpeed))
		l.SetField(tbl, "sink", lua.LNumber(val.Sink))
		l.SetField(tbl, "smoothing", lua.LNumber(val.Smoothing))
		l.SetField(tbl, "walk_speed", lua.LNumber(val.WalkSpeed))
	case *mt.ToCltOverrideDayNightRatio:
		l.SetField(tbl, "override", lua.LBool(val.Override))
		l.SetField(tbl, "ratio", lua.LNumber(val.Ratio))
	case *mt.ToCltPlaySound:
		l.SetField(tbl, "ephemeral", lua.LBool(val.Ephemeral))
		l.SetField(tbl, "fade", lua.LNumber(val.Fade))
		l.SetField(tbl, "gain", lua.LNumber(val.Gain))
		l.SetField(tbl, "id", lua.LNumber(val.ID))
		l.SetField(tbl, "loop", lua.LBool(val.Loop))
		l.SetField(tbl, "name", lua.LString(string(val.Name)))
		l.SetField(tbl, "pitch", lua.LNumber(val.Pitch))
		l.SetField(tbl, "pos", Vec3(l, [3]lua.LNumber{lua.LNumber(val.Pos[0]), lua.LNumber(val.Pos[1]), lua.LNumber(val.Pos[2])}))
		l.SetField(tbl, "src_aoid", lua.LNumber(val.SrcAOID))
		l.SetField(tbl, "src_type", SoundSrcType(l, val.SrcType))
	case *mt.ToCltPrivs:
		l.SetField(tbl, "privs", StringSet(l, val.Privs))
	case *mt.ToCltRemoveNode:
		l.SetField(tbl, "pos", Vec3(l, [3]lua.LNumber{lua.LNumber(val.Pos[0]), lua.LNumber(val.Pos[1]), lua.LNumber(val.Pos[2])}))
	case *mt.ToCltRmHUD:
		l.SetField(tbl, "id", lua.LNumber(val.ID))
	case *mt.ToCltSetHotbarParam:
		l.SetField(tbl, "img", lua.LString(string(val.Img)))
		l.SetField(tbl, "param", HotbarParam(l, val.Param))
		l.SetField(tbl, "size", lua.LNumber(val.Size))
	case *mt.ToCltShowFormspec:
		l.SetField(tbl, "formname", lua.LString(string(val.Formname)))
		l.SetField(tbl, "formspec", lua.LString(string(val.Formspec)))
	case *mt.ToCltSkyParams:
		l.SetField(tbl, "bg_color", Color(l, val.BgColor))
		l.SetField(tbl, "clouds", lua.LBool(val.Clouds))
		if val.Type == "regular" {
			l.SetField(tbl, "dawn_horizon", Color(l, val.DawnHorizon))
		}
		if val.Type == "regular" {
			l.SetField(tbl, "dawn_sky", Color(l, val.DawnSky))
		}
		if val.Type == "regular" {
			l.SetField(tbl, "day_horizon", Color(l, val.DayHorizon))
		}
		if val.Type == "regular" {
			l.SetField(tbl, "day_sky", Color(l, val.DaySky))
		}
		l.SetField(tbl, "fog_tint_type", lua.LString(string(val.FogTintType)))
		if val.Type == "regular" {
			l.SetField(tbl, "indoor", Color(l, val.Indoor))
		}
		l.SetField(tbl, "moon_fog_tint", Color(l, val.MoonFogTint))
		if val.Type == "regular" {
			l.SetField(tbl, "night_horizon", Color(l, val.NightHorizon))
		}
		if val.Type == "regular" {
			l.SetField(tbl, "night_sky", Color(l, val.NightSky))
		}
		l.SetField(tbl, "sun_fog_tint", Color(l, val.SunFogTint))
		if val.Type == "skybox" {
			l.SetField(tbl, "textures", TextureList(l, val.Textures))
		}
		l.SetField(tbl, "type", lua.LString(string(val.Type)))
	case *mt.ToCltSpawnParticle:
		l.SetField(tbl, "acc", Vec3(l, [3]lua.LNumber{lua.LNumber(val.Acc[0]), lua.LNumber(val.Acc[1]), lua.LNumber(val.Acc[2])}))
		l.SetField(tbl, "anim_params", TileAnim(l, val.AnimParams))
		l.SetField(tbl, "ao_collision", lua.LBool(val.AOCollision))
		l.SetField(tbl, "collide", lua.LBool(val.Collide))
		l.SetField(tbl, "collision_rm", lua.LBool(val.CollisionRm))
		l.SetField(tbl, "expiration_time", lua.LNumber(val.ExpirationTime))
		l.SetField(tbl, "glow", lua.LNumber(val.Glow))
		l.SetField(tbl, "node_param0", lua.LNumber(val.NodeParam0))
		l.SetField(tbl, "node_param2", lua.LNumber(val.NodeParam2))
		l.SetField(tbl, "node_tile", lua.LNumber(val.NodeTile))
		l.SetField(tbl, "pos", Vec3(l, [3]lua.LNumber{lua.LNumber(val.Pos[0]), lua.LNumber(val.Pos[1]), lua.LNumber(val.Pos[2])}))
		l.SetField(tbl, "size", lua.LNumber(val.Size))
		l.SetField(tbl, "texture", lua.LString(string(val.Texture)))
		l.SetField(tbl, "vel", Vec3(l, [3]lua.LNumber{lua.LNumber(val.Vel[0]), lua.LNumber(val.Vel[1]), lua.LNumber(val.Vel[2])}))
		l.SetField(tbl, "vertical", lua.LBool(val.Vertical))
	case *mt.ToCltSRPBytesSaltB:
		l.SetField(tbl, "b", lua.LString(string(val.B)))
		l.SetField(tbl, "salt", lua.LString(string(val.Salt)))
	case *mt.ToCltStarParams:
		l.SetField(tbl, "color", Color(l, val.Color))
		l.SetField(tbl, "count", lua.LNumber(val.Count))
		l.SetField(tbl, "size", lua.LNumber(val.Size))
		l.SetField(tbl, "visible", lua.LBool(val.Visible))
	case *mt.ToCltStopSound:
		l.SetField(tbl, "id", lua.LNumber(val.ID))
	case *mt.ToCltSunParams:
		l.SetField(tbl, "rise", lua.LString(string(val.Rise)))
		l.SetField(tbl, "rising", lua.LBool(val.Rising))
		l.SetField(tbl, "size", lua.LNumber(val.Size))
		l.SetField(tbl, "texture", lua.LString(string(val.Texture)))
		l.SetField(tbl, "tone_map", lua.LString(string(val.ToneMap)))
		l.SetField(tbl, "visible", lua.LBool(val.Visible))
	case *mt.ToCltTimeOfDay:
		l.SetField(tbl, "speed", lua.LNumber(val.Speed))
		l.SetField(tbl, "time", lua.LNumber(val.Time))
	case *mt.ToCltUpdatePlayerList:
		l.SetField(tbl, "players", StringList(l, val.Players))
		l.SetField(tbl, "type", PlayerListUpdateType(l, val.Type))
	}
	return tbl
}
