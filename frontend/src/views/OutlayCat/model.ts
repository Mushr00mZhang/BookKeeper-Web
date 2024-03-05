import * as PagedList from '@/utils/PagedList';
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
}
