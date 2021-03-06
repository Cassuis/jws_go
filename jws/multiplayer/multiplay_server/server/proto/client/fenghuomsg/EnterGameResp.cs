// automatically generated by the FlatBuffers compiler, do not modify

namespace fenghuomsg
{

using System;
using FlatBuffers;

/// [RPC]进入同步战斗服务器 resp
public sealed class EnterGameResp : Table {
  public static EnterGameResp GetRootAsEnterGameResp(ByteBuffer _bb) { return GetRootAsEnterGameResp(_bb, new EnterGameResp()); }
  public static EnterGameResp GetRootAsEnterGameResp(ByteBuffer _bb, EnterGameResp obj) { return (obj.__init(_bb.GetInt(_bb.Position) + _bb.Position, _bb)); }
  public EnterGameResp __init(int _i, ByteBuffer _bb) { bb_pos = _i; bb = _bb; return this; }

  ///当前玩家索引
  public int Myidx { get { int o = __offset(4); return o != 0 ? bb.GetInt(o + bb_pos) : (int)0; } }
  ///当前服务器状态
  public fenghuomsg.FenghuoGameStatus Gamestatus { get { int o = __offset(6); return o != 0 ? (fenghuomsg.FenghuoGameStatus)bb.GetSbyte(o + bb_pos) : fenghuomsg.FenghuoGameStatus.WaitingInit; } }
  ///敌兵所有血量
  public int GetEnemiesHp(int j) { int o = __offset(8); return o != 0 ? bb.GetInt(__vector(o) + j * 4) : (int)0; }
  public int EnemiesHpLength { get { int o = __offset(8); return o != 0 ? __vector_len(o) : 0; } }
  public ArraySegment<byte>? GetEnemiesHpBytes() { return __vector_as_arraysegment(8); }
  ///玩家血量
  public int GetHps(int j) { int o = __offset(10); return o != 0 ? bb.GetInt(__vector(o) + j * 4) : (int)0; }
  public int HpsLength { get { int o = __offset(10); return o != 0 ? __vector_len(o) : 0; } }
  public ArraySegment<byte>? GetHpsBytes() { return __vector_as_arraysegment(10); }
  ///玩家Avatar属性
  public fenghuomsg.AccountInfo GetAvatars(int j) { return GetAvatars(new fenghuomsg.AccountInfo(), j); }
  public fenghuomsg.AccountInfo GetAvatars(fenghuomsg.AccountInfo obj, int j) { int o = __offset(12); return o != 0 ? obj.__init(__indirect(__vector(o) + j * 4), bb) : null; }
  public int AvatarsLength { get { int o = __offset(12); return o != 0 ? __vector_len(o) : 0; } }

  public static Offset<EnterGameResp> CreateEnterGameResp(FlatBufferBuilder builder,
      int myidx = 0,
      fenghuomsg.FenghuoGameStatus gamestatus = fenghuomsg.FenghuoGameStatus.WaitingInit,
      VectorOffset enemiesHpOffset = default(VectorOffset),
      VectorOffset HpsOffset = default(VectorOffset),
      VectorOffset avatarsOffset = default(VectorOffset)) {
    builder.StartObject(5);
    EnterGameResp.AddAvatars(builder, avatarsOffset);
    EnterGameResp.AddHps(builder, HpsOffset);
    EnterGameResp.AddEnemiesHp(builder, enemiesHpOffset);
    EnterGameResp.AddMyidx(builder, myidx);
    EnterGameResp.AddGamestatus(builder, gamestatus);
    return EnterGameResp.EndEnterGameResp(builder);
  }

  public static void StartEnterGameResp(FlatBufferBuilder builder) { builder.StartObject(5); }
  public static void AddMyidx(FlatBufferBuilder builder, int myidx) { builder.AddInt(0, myidx, 0); }
  public static void AddGamestatus(FlatBufferBuilder builder, fenghuomsg.FenghuoGameStatus gamestatus) { builder.AddSbyte(1, (sbyte)gamestatus, 0); }
  public static void AddEnemiesHp(FlatBufferBuilder builder, VectorOffset enemiesHpOffset) { builder.AddOffset(2, enemiesHpOffset.Value, 0); }
  public static VectorOffset CreateEnemiesHpVector(FlatBufferBuilder builder, int[] data) { builder.StartVector(4, data.Length, 4); for (int i = data.Length - 1; i >= 0; i--) builder.AddInt(data[i]); return builder.EndVector(); }
  public static void StartEnemiesHpVector(FlatBufferBuilder builder, int numElems) { builder.StartVector(4, numElems, 4); }
  public static void AddHps(FlatBufferBuilder builder, VectorOffset HpsOffset) { builder.AddOffset(3, HpsOffset.Value, 0); }
  public static VectorOffset CreateHpsVector(FlatBufferBuilder builder, int[] data) { builder.StartVector(4, data.Length, 4); for (int i = data.Length - 1; i >= 0; i--) builder.AddInt(data[i]); return builder.EndVector(); }
  public static void StartHpsVector(FlatBufferBuilder builder, int numElems) { builder.StartVector(4, numElems, 4); }
  public static void AddAvatars(FlatBufferBuilder builder, VectorOffset avatarsOffset) { builder.AddOffset(4, avatarsOffset.Value, 0); }
  public static VectorOffset CreateAvatarsVector(FlatBufferBuilder builder, Offset<fenghuomsg.AccountInfo>[] data) { builder.StartVector(4, data.Length, 4); for (int i = data.Length - 1; i >= 0; i--) builder.AddOffset(data[i].Value); return builder.EndVector(); }
  public static void StartAvatarsVector(FlatBufferBuilder builder, int numElems) { builder.StartVector(4, numElems, 4); }
  public static Offset<EnterGameResp> EndEnterGameResp(FlatBufferBuilder builder) {
    int o = builder.EndObject();
    return new Offset<EnterGameResp>(o);
  }
};


}
