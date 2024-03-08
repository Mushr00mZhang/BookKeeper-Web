import * as PagedList from '@/utils/PagedList';
import { IResult, Result } from '@/utils/Result';
import axios from 'axios';
/**
 * 支出类型基础
 */
export interface IBase {
  /**
   * 父级Id
   */
  parentId: string;
  /**
   * 名称
   */
  name: string;
  /**
   * 单位
   */
  unit: string;
  /**
   * 排序
   */
  sort: number;
  /**
   * 是否固定
   */
  stable: boolean;
  /**
   * 备注
   */
  remark: string;
}
/**
 * 支出类型Dto
 */
export interface IDto extends IBase {
  /**
   * 主键id
   */
  id: string;
  /**
   * 子集列表
   */
  children: IDto[];
  // outlays:IDto[]
}
/**
 * 支出类型创建Dto
 */
export interface ICreateDto extends IBase {}
/**
 * 支出类型更新Dto
 */
export interface IUpdateDto extends IBase {
  /**
   * 主键id
   */
  id: string;
}
/**
 * 支出类型列表Dto
 */
export interface IListDto {
  parentId?: string | null;
}
/**
 * 支出类型分页列表Dto
 */
export interface IPagedListDto extends IListDto, PagedList.IGetDto {}
/**
 * 支出类型
 */
export class OutlayCat implements IDto {
  id: string = '';
  parentId: string = '';
  name: string = '';
  unit: string = '';
  sort: number = 0;
  stable: boolean = false;
  remark: string = '';
  children: OutlayCat[] = [];
  constructor(dto: IDto) {
    this.id = dto.id;
    this.parentId = dto.parentId;
    this.name = dto.name;
    this.unit = dto.unit;
    this.sort = dto.sort;
    this.stable = dto.stable;
    this.remark = dto.remark;
    this.children = dto.children.map((i) => new OutlayCat(i));
  }
  static async list(dto: IListDto) {
    const url = `/api/outlaycats`;
    const res = await axios.get<Result<IDto[]>>(url, { params: dto });
    return (res.data.result || []).map((i) => new OutlayCat(i));
  }
  static async pagedlist(dto: IPagedListDto) {
    const url = `/api/outlaycats`;
    const res = await axios.get<Result<PagedList.IDto<IDto>>>(url, { params: dto });
    const items = (res.data.result?.items || []).map((i) => new OutlayCat(i));
    return {
      total: res.data.result?.total || 0,
      items,
    } as PagedList.IDto<IDto>;
  }
  static async get(id: string) {
    const url = `/api/outlaycats/${id}`;
    const res = await axios.get<Result<IDto>>(url);
    return (res.data.result && new OutlayCat(res.data.result)) || null;
  }
  static async create(dto: ICreateDto) {
    const url = `/api/outlaycats`;
    const res = await axios.put<Result<string>>(url, dto);
    return res.data.result;
  }
  static async update(dto: IUpdateDto) {
    const url = `/api/outlaycats/${dto.id}`;
    const res = await axios.post<Result<boolean>>(url, dto);
    return res.data.result;
  }
  static async delete(id: string) {
    const url = `/api/outlaycats/${id}`;
    const res = await axios.delete<Result<boolean>>(url);
    return res.data.result;
  }
  readonly update = () => OutlayCat.update(this);
  readonly delete = () => OutlayCat.delete(this.id);
}
