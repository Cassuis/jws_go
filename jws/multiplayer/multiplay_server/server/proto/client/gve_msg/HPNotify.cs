// automatically generated by the FlatBuffers compiler, do not modify

namespace gve_msg
{

using System;
using FlatBuffers;

/// [Notify]伤害\损失HP通知
public sealed class HPNotify : Table {
  public static HPNotify GetRootAsHPNotify(ByteBuffer _bb) { return GetRootAsHPNotify(_bb, new HPNotify()); }
  public static HPNotify GetRootAsHPNotify(ByteBuffer _bb, HPNotify obj) { return (obj.__init(_bb.GetInt(_bb.Position) + _bb.Position, _bb)); }
  public HPNotify __init(int _i, ByteBuffer _bb) { bb_pos = _i; bb = _bb; return this; }

  ///玩家账号ID
  public string AccountId { get { int o = __offset(4); return o != 0 ? __string(o + bb_pos) : null; } }
  public ArraySegment<byte>? GetAccountIdBytes() { return __vector_as_arraysegment(4); }
  ///请求进入房间ID
  public string RoomID { get { int o = __offset(6); return o != 0 ? __string(o + bb_pos) : null; } }
  public ArraySegment<byte>? GetRoomIDBytes() { return __vector_as_arraysegment(6); }
  ///玩家自身血量变化
  public int PlayerHpD { get { int o = __offset(8); return o != 0 ? bb.GetInt(o + bb_pos) : (int)0; } }
  ///玩家对Boss造成的血量变化
  public int GetBossHpD(int j) { int o = __offset(10); return o != 0 ? bb.GetInt(__vector(o) + j * 4) : (int)0; }
  public int BossHpDLength { get { int o = __offset(10); return o != 0 ? __vector_len(o) : 0; } }
  public ArraySegment<byte>? GetBossHpDBytes() { return __vector_as_arraysegment(10); }

  public static Offset<HPNotify> CreateHPNotify(FlatBufferBuilder builder,
      StringOffset accountIdOffset = default(StringOffset),
      StringOffset roomIDOffset = default(StringOffset),
      int playerHpD = 0,
      VectorOffset bossHpDOffset = default(VectorOffset)) {
    builder.StartObject(4);
    HPNotify.AddBossHpD(builder, bossHpDOffset);
    HPNotify.AddPlayerHpD(builder, playerHpD);
    HPNotify.AddRoomID(builder, roomIDOffset);
    HPNotify.AddAccountId(builder, accountIdOffset);
    return HPNotify.EndHPNotify(builder);
  }

  public static void StartHPNotify(FlatBufferBuilder builder) { builder.StartObject(4); }
  public static void AddAccountId(FlatBufferBuilder builder, StringOffset accountIdOffset) { builder.AddOffset(0, accountIdOffset.Value, 0); }
  public static void AddRoomID(FlatBufferBuilder builder, StringOffset roomIDOffset) { builder.AddOffset(1, roomIDOffset.Value, 0); }
  public static void AddPlayerHpD(FlatBufferBuilder builder, int playerHpD) { builder.AddInt(2, playerHpD, 0); }
  public static void AddBossHpD(FlatBufferBuilder builder, VectorOffset bossHpDOffset) { builder.AddOffset(3, bossHpDOffset.Value, 0); }
  public static VectorOffset CreateBossHpDVector(FlatBufferBuilder builder, int[] data) { builder.StartVector(4, data.Length, 4); for (int i = data.Length - 1; i >= 0; i--) builder.AddInt(data[i]); return builder.EndVector(); }
  public static void StartBossHpDVector(FlatBufferBuilder builder, int numElems) { builder.StartVector(4, numElems, 4); }
  public static Offset<HPNotify> EndHPNotify(FlatBufferBuilder builder) {
    int o = builder.EndObject();
    return new Offset<HPNotify>(o);
  }
  public static void FinishHPNotifyBuffer(FlatBufferBuilder builder, Offset<HPNotify> offset) { builder.Finish(offset.Value); }
};


}