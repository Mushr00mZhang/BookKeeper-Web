import * as PagedList from '@/utils/PagedList';
import { Result } from '@/utils/Result';
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
  /** 主键id */
  id: string | null;
  /** 是否含有子集 */
  hasChildren: boolean;
  // /**
  //  * 子集列表
  //  */
  // children: IDto[] | null;
  // outlays:IDto[]
}
/** 支出类型创建Dto */
export interface ICreateDto extends IBase {}
/** 支出类型更新Dto */
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
  id: string | null = null;
  parentId: string = ROOT_ID;
  name: string = '';
  unit: string = '';
  sort: number = 0;
  stable: boolean = false;
  remark: string = '';
  hasChildren: boolean = false;
  children: OutlayCat[] = [];
  constructor(dto: IDto) {
    this.id = dto.id;
    this.parentId = dto.parentId;
    this.name = dto.name;
    this.unit = dto.unit;
    this.sort = dto.sort;
    this.stable = dto.stable;
    this.remark = dto.remark;
    this.hasChildren = dto.hasChildren;
    // this.children = dto?.children?.map((i) => new OutlayCat(i)) || [];
  }
  static readonly list = async (dto: IListDto) => {
    const url = `/api/outlaycats`;
    const res = await axios.get<Result<IDto[]>>(url, { params: dto });
    return (res.data.result || []).map((i) => new OutlayCat(i));
  };
  static readonly pagedlist = async (dto: IPagedListDto) => {
    const url = `/api/outlaycats`;
    const res = await axios.get<Result<PagedList.IDto<IDto>>>(url, { params: dto });
    const items = (res.data.result?.items || []).map((i) => new OutlayCat(i));
    return {
      total: res.data.result?.total || 0,
      items,
    } as PagedList.IDto<IDto>;
  };
  static readonly get = async (id: string) => {
    const url = `/api/outlaycats/${id}`;
    const res = await axios.get<Result<IDto>>(url);
    return (res.data.result && new OutlayCat(res.data.result)) || null;
  };
  static readonly create = async (dto: ICreateDto) => {
    const url = `/api/outlaycats`;
    const res = await axios.post<Result<string>>(url, dto);
    return res.data.result;
  };
  static readonly update = async (dto: IUpdateDto) => {
    const url = `/api/outlaycats/${dto.id}`;
    const res = await axios.put<Result<boolean>>(url, dto);
    return res.data.result;
  };
  static readonly delete = async (id: string) => {
    const url = `/api/outlaycats/${id}`;
    const res = await axios.delete<Result<boolean>>(url);
    return res.data.result;
  };
  static readonly getParentIds = async (parentId: string): Promise<string[]> => {
    if (parentId === ROOT_ID) return [parentId];
    const cat = await OutlayCat.get(parentId);
    if (cat) return [parentId, ...(await OutlayCat.getParentIds(cat.parentId))];
    return [parentId];
  };
  readonly create = async () => {
    const id = await OutlayCat.create(this as ICreateDto);
    if (id) this.id = id;
    return id;
  };
  readonly update = () => {
    if (this.id) return OutlayCat.update(this as IUpdateDto);
    return false;
  };
  readonly delete = () => {
    if (this.id) return OutlayCat.delete(this.id);
    return false;
  };
  readonly getParentIds = async () => OutlayCat.getParentIds(this.parentId);
}
export const ROOT_ID = '00000000-0000-0000-0000-000000000000';
export const ROOT = new OutlayCat({
  id: ROOT_ID,
  parentId: '',
  name: '根目录',
  unit: '',
  sort: 0,
  stable: false,
  remark: '',
  hasChildren: true,
  // children: [],
});
