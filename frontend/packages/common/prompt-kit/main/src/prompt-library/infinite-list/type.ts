/*
 * Copyright 2025 coze-dev Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
 
import { type ReactElement, type RefObject } from 'react';

import {
  type ListProps,
  type EmptyStateProps,
  type ButtonProps,
} from '@coze-arch/coze-design';

export interface EmptyProps {
  className?: string;
  isError?: boolean;
  isLoading?: boolean;
  loadRetry?: () => void; //重试加载
  size?: EmptyStateProps['size'];
  text?: {
    emptyTitle?: string;
    emptyDesc?: string;
    searchEmptyTitle?: string;
  };
  btn?: {
    emptyClick?: () => void; //
    emptyText?: string;
    emptyButtonProps?: ButtonProps;
  };
  icon?: ReactElement;

  renderEmpty?: (
    emptyProps: Omit<EmptyProps, 'renderEmpty'>,
  ) => React.ReactNode | null;
}
export interface FooterProps {
  isError?: boolean; // 是否加载出错
  isLoading?: boolean; // 是否加载中
  noMore?: boolean; //没有更多数据
  isNeedBtnLoadMore?: boolean;
  dataNum?: number;
  loadRetry?: () => void; //重试加载
  renderFooter?: (
    footerProps: Omit<FooterProps, 'renderFooter'>,
  ) => React.ReactNode | null;
}

export interface InfiniteListDataProps<T> {
  list: T[];
  hasMore?: boolean;
  // nextPage: number;
  cursor: string;
  [key: string]: unknown;
}

export interface ScrollProps<T> {
  threshold?: number; //距离下方多长距离，开始加载数据
  targetRef?: RefObject<HTMLDivElement>; // 监听滚动的Dom 引用
  loadData: (
    current: InfiniteListDataProps<T>,
  ) => Promise<InfiniteListDataProps<T>>; // 加载更多数据
  reloadDeps?: unknown[]; // 重新加载数据依赖
  isNeedBtnLoadMore?: boolean;
  isLoading?: boolean; // 是否加载中
  resetDataIfReload?: boolean; // 当reload时，是否先reset列表已存在数据，默认为true
}

export interface InfiniteListProps<T>
  extends Partial<
    Pick<ListProps<T>, 'className' | 'emptyContent' | 'grid' | 'renderItem'>
  > {
  containerClassName?: string;
  canShowData?: boolean; //是否能够显示数据了
  isSearching?: boolean; // 是否搜索中，主要是用于错误显示的时候，选择文案使用
  itemClassName?: string | ((item: T) => string);
  isNeedBtnLoadMore?: boolean;
  isResponsive?: boolean;
  emptyConf?: {
    className?: string;
    renderEmpty?: EmptyProps['renderEmpty'];
    text?: EmptyProps['text'];
    btn?: EmptyProps['btn'];
    icon?: EmptyProps['icon'];
    size?: EmptyProps['size'];
  };
  renderFooter?: FooterProps['renderFooter'];
  scrollConf: ScrollProps<T>;
  rowKey?: string;
  retryFunc?: () => void;
  onChangeState?: (loading: boolean, data: T[]) => void;
}

export interface InfiniteListRef<T> {
  reload: () => void;
  getDataList: () => T[]; // 获取当前列表数据
}
