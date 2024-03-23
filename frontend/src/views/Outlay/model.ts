import * as PagedList from '@/utils/PagedList';
import { Result } from '@/utils/Result';
import axios from 'axios';
import { IModel as OutlayCatModel } from '@/views/OutlayCat/model';
/** 支出基础 */
export interface IBase {
  /** 名称 */
  name: string;
  /** 类型Id */
  catId: string;
  /** 金额 */
  money: number;
  /** 原价 */
  original: number;
  /** 数量 */
  amount: number;
  /** 单位 */
  unit: string;
  /** 时间 */
  time: string;
  /** 用户Id */
  userId: string;
}
/** 支出模型 */
export interface IModel extends IBase {
  /** 主键id */
  id?: string;
}
/** 支出Dto */
export interface IDto extends IModel {
  /** 类型 */
  cat: OutlayCatModel;
}
/** 支出创建Dto */
export interface ICreateDto extends IModel {
  id: undefined;
}
/** 支出更新Dto */
export interface IUpdateDto extends IModel {
  id: string;
}
/** 支出列表Dto */
export interface IListDto {
  /** 类型Id */
  catId?: string;
  /** 最低金额 */
  lowMoney?: number;
  /** 最高金额 */
  topMoney?: number;
  /** 开始时间 */
  sTime?: string;
  /** 结束时间 */
  eTime?: string;
  /** 用户Id */
  userId?: string;
}
/** 支出列表Dto */
export interface IPagedListDto extends IListDto, PagedList.IGetDto {}
/**
 * 支出
 */
export class Outlay implements IDto {
  id?: string;
  name: string;
  catId: string = '';
  money: number = 0;
  original: number = 0;
  amount: number = 0;
  unit: string = '';
  time: string = '';
  userId: string = '';
  cat: OutlayCatModel;
  constructor(dto: IDto) {
    this.id = dto.id;
    this.name = dto.name;
    this.catId = dto.catId;
    this.money = dto.money;
    this.original = dto.original;
    this.amount = dto.amount;
    this.unit = dto.unit;
    this.time = dto.time;
    this.userId = dto.userId;
    this.cat = dto.cat;
  }
  static readonly list = async (dto: IListDto) => {
    const url = `/api/outlays`;
    const res = await axios.get<Result<IDto[]>>(url, { params: dto });
    return (res.data.result || []).map((i) => new Outlay(i));
  };
  static readonly pagedlist = async (dto: IPagedListDto) => {
    const url = `/api/outlays`;
    const res = await axios.get<Result<PagedList.IDto<IDto>>>(url, { params: dto });
    const items = (res.data.result?.items || []).map((i) => new Outlay(i));
    return {
      total: res.data.result?.total || 0,
      items,
    } as PagedList.IDto<IDto>;
  };
  static readonly get = async (id: string) => {
    const url = `/api/outlays/${id}`;
    const res = await axios.get<Result<IDto>>(url);
    return (res.data.result && new Outlay(res.data.result)) || null;
  };
  static readonly create = async (dto: ICreateDto) => {
    const url = `/api/outlays`;
    const res = await axios.post<Result<string>>(url, dto);
    return res.data.result;
  };
  static readonly update = async (dto: IUpdateDto) => {
    const url = `/api/outlays/${dto.id}`;
    const res = await axios.put<Result<boolean>>(url, dto);
    return res.data.result;
  };
  static readonly delete = async (id: string) => {
    const url = `/api/outlays/${id}`;
    const res = await axios.delete<Result<boolean>>(url);
    return res.data.result;
  };
  readonly create = async () => {
    const id = await Outlay.create(this as ICreateDto);
    if (id) this.id = id;
    return id;
  };
  readonly update = () => {
    if (this.id) return Outlay.update(this as IUpdateDto);
    return false;
  };
  readonly delete = () => {
    if (this.id) return Outlay.delete(this.id);
    return false;
  };
}
