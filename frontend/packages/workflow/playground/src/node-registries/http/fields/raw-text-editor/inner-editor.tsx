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
 
import { useRef, type FC } from 'react';

import classNames from 'classnames';
import { type ExpressionEditorTreeNode } from '@coze-workflow/components';
import {
  useAvailableWorkflowVariables,
  useGetWorkflowVariableByKeyPath,
  isGlobalVariableKey,
} from '@coze-workflow/variable';

import { VariableExtension } from '@/node-registries/http/components/variable-support';
import { useNodeServiceAndRefreshForTitleChange } from '@/form-extensions/hooks/use-node-available-variables';
import {
  formatDataWithGlobalVariable,
  processDataSourceLabelRender,
} from '@/form-extensions/components/tree-variable-selector/utils';
import { useFormatVariableDataSource } from '@/form-extensions/components/tree-variable-selector/useFormatVariableDataSource';
import useGlobalVariableCache from '@/form-extensions/components/tree-variable-selector/use-global-variable-cache';
import GlobalVarIcon from '@/form-extensions/components/tree-variable-selector/global-var-icon';
import { NodeIcon } from '@/components/node-icon';

import { useVariableWithNodeInfo } from '../hooks/use-variable-with-node';
import { useVariableTree } from '../hooks/use-variable-tree';
import { BaseRawTextEditor } from '../../components/base-editor';

import styles from './index.module.less';

interface RawTextEditorContainerProps {
  name: string;
  value: string;
  onChange: (value: string) => void;
  key?: string;
  placeholder?: string | HTMLElement;
  readonly?: boolean;
  disableSuggestion?: boolean;
  disableCounter?: boolean;
  minRows?: number;
  maxLength?: number;
  onBlur?: () => void;
  onFocus?: () => void;
  isError?: boolean;
  minHeight?: string | number;
}

/**
 * 全局变量提示逻辑层
 */
export const InnerEditor: FC<RawTextEditorContainerProps> = props => {
  const { value, onChange, placeholder, minHeight, readonly = false } = props;

  const editorContainerRef = useRef<HTMLDivElement>(null);

  const getVariableByKeyPath = useGetWorkflowVariableByKeyPath();
  const allVariables = useAvailableWorkflowVariables();
  const { getNodeInfoInVariableMeta } =
    useNodeServiceAndRefreshForTitleChange();
  const availableVariables = useVariableWithNodeInfo(
    allVariables,
    getNodeInfoInVariableMeta,
  );
  /** 默认兜底的可选择变量数据 */
  const defaultVariableDataSource = useFormatVariableDataSource({
    disabledTypes: [],
  });

  const dataSourceWithGlobal = useGlobalVariableCache(
    defaultVariableDataSource,
  );

  // 处理 DataSource 数据，添加部分字段 / 渲染
  const variableDataSource = processDataSourceLabelRender({
    dataSource: dataSourceWithGlobal,
    icon: node =>
      isGlobalVariableKey(node.value) ? (
        <GlobalVarIcon nodeId={node.value} />
      ) : (
        <NodeIcon
          size={16}
          alt="logo"
          nodeId={node.value}
          className="leading-[0px]"
        />
      ),
  });

  const variableDataSourceWithGroup =
    formatDataWithGlobalVariable(variableDataSource);

  const variableTree: ExpressionEditorTreeNode[] = useVariableTree({
    variables: allVariables,
    getNodeInfoInVariableMeta,
  });

  return (
    <div className="w-full">
      <BaseRawTextEditor
        ref={editorContainerRef}
        readonly={readonly}
        value={value}
        className={classNames(
          styles['editor-render'],
          styles['editor-render-cm-content'],
        )}
        placeholder={placeholder}
        // dataTestID={dataTestID}
        onChange={onChange}
        minHeight={minHeight}
      />
      <VariableExtension
        readonly={readonly}
        availableVariables={availableVariables}
        variableDataSource={variableDataSourceWithGroup}
        getVariableByKeyPath={getVariableByKeyPath}
        variableTree={variableTree}
      />
    </div>
  );
};
