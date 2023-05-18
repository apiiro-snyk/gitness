import React, { useMemo } from 'react'
import {
  ButtonSize,
  ButtonVariation,
  Color,
  Container,
  FlexExpander,
  Heading,
  Layout,
  Tabs,
  Utils
} from '@harness/uicore'
import { Render } from 'react-jsx-match'
import { useHistory } from 'react-router-dom'
import { SourceCodeViewer } from 'components/SourceCodeViewer/SourceCodeViewer'
import type { OpenapiContentInfo, RepoFileContent } from 'services/code'
import {
  decodeGitContent,
  findMarkdownInfo,
  GitCommitAction,
  GitInfoProps,
  isRefATag,
  makeDiffRefs
} from 'utils/GitUtils'
import { filenameToLanguage } from 'utils/Utils'
import { useAppContext } from 'AppContext'
import { LatestCommitForFile } from 'components/LatestCommit/LatestCommit'
import { useCommitModal } from 'components/CommitModalButton/CommitModalButton'
import { useStrings } from 'framework/strings'
import { OptionsMenuButton } from 'components/OptionsMenuButton/OptionsMenuButton'
import { PlainButton } from 'components/PlainButton/PlainButton'
import { Readme } from '../FolderContent/Readme'
import { GitBlame } from './GitBlame'
import css from './FileContent.module.scss'

enum FileSection {
  CONTENT = 'content',
  BLAME = 'blame',
  HISTORY = 'history'
}

export function FileContent({
  repoMetadata,
  gitRef,
  resourcePath,
  resourceContent
}: Pick<GitInfoProps, 'repoMetadata' | 'gitRef' | 'resourcePath' | 'resourceContent'>) {
  const { routes } = useAppContext()
  const { getString } = useStrings()
  const history = useHistory()
  const content = useMemo(
    () => decodeGitContent((resourceContent?.content as RepoFileContent)?.data),
    [resourceContent?.content]
  )
  const markdownInfo = useMemo(() => findMarkdownInfo(resourceContent), [resourceContent])
  const [openDeleteFileModal] = useCommitModal({
    repoMetadata,
    gitRef,
    resourcePath,
    commitAction: GitCommitAction.DELETE,
    commitTitlePlaceHolder: getString('deleteFile').replace('__path__', resourcePath),
    onSuccess: (_commitInfo, newBranch) => {
      if (newBranch) {
        history.replace(
          routes.toCODECompare({
            repoPath: repoMetadata.path as string,
            diffRefs: makeDiffRefs(repoMetadata?.default_branch as string, newBranch)
          })
        )
      } else {
        history.push(
          routes.toCODERepository({
            repoPath: repoMetadata.path as string,
            gitRef
          })
        )
      }
    }
  })

  return (
    <Container className={css.tabsContainer}>
      <Tabs
        id="fileTabs"
        defaultSelectedTabId={FileSection.CONTENT}
        large={false}
        tabList={[
          {
            id: FileSection.CONTENT,
            title: getString('content'),
            panel: (
              <Container className={css.fileContent}>
                <Layout.Vertical spacing="small">
                  <LatestCommitForFile
                    repoMetadata={repoMetadata}
                    latestCommit={resourceContent.latest_commit}
                    standaloneStyle
                  />
                  <Container className={css.container} background={Color.WHITE}>
                    <Layout.Horizontal padding="small" className={css.heading}>
                      <Heading level={5} color={Color.BLACK}>
                        {resourceContent.name}
                      </Heading>
                      <FlexExpander />
                      <Layout.Horizontal spacing="xsmall" style={{ alignItems: 'center' }}>
                        <PlainButton
                          withoutCurrentColor
                          size={ButtonSize.SMALL}
                          variation={ButtonVariation.TERTIARY}
                          iconProps={{ size: 16 }}
                          text={getString('edit')}
                          icon="code-edit"
                          tooltip={isRefATag(gitRef) ? getString('editNotAllowed') : undefined}
                          tooltipProps={{ isDark: true }}
                          disabled={isRefATag(gitRef)}
                          onClick={() => {
                            history.push(
                              routes.toCODEFileEdit({
                                repoPath: repoMetadata.path as string,
                                gitRef,
                                resourcePath
                              })
                            )
                          }}
                        />
                        <OptionsMenuButton
                          isDark={true}
                          icon="Options"
                          iconProps={{ size: 14 }}
                          style={{ padding: '5px' }}
                          width="145px"
                          items={[
                            {
                              hasIcon: true,
                              iconName: 'arrow-right',
                              text: getString('viewRaw'),
                              onClick: () => {
                                window.open(
                                  `/code/api/v1/repos/${
                                    repoMetadata?.path
                                  }/+/raw/${resourcePath}?${`git_ref=${gitRef}`}`,
                                  '_blank'
                                )
                              }
                            },
                            '-',
                            {
                              hasIcon: true,
                              iconName: 'cloud-download',
                              text: getString('download'),
                              download: resourceContent?.name || 'download',
                              href: `/code/api/v1/repos/${
                                repoMetadata?.path
                              }/+/raw/${resourcePath}?${`git_ref=${gitRef}`}`
                            },
                            {
                              hasIcon: true,
                              iconName: 'code-copy',
                              iconSize: 16,
                              text: getString('copy'),
                              onClick: () => Utils.copy(content)
                            },
                            {
                              hasIcon: true,
                              iconName: 'code-delete',
                              iconSize: 16,
                              title: getString(isRefATag(gitRef) ? 'deleteNotAllowed' : 'delete'),
                              disabled: isRefATag(gitRef),
                              text: getString('delete'),
                              onClick: openDeleteFileModal
                            }
                          ]}
                        />
                      </Layout.Horizontal>
                    </Layout.Horizontal>

                    <Render when={(resourceContent?.content as RepoFileContent)?.data}>
                      <Container className={css.content}>
                        <Render when={!markdownInfo}>
                          <SourceCodeViewer language={filenameToLanguage(resourceContent?.name)} source={content} />
                        </Render>
                        <Render when={markdownInfo}>
                          <Readme
                            metadata={repoMetadata}
                            readmeInfo={markdownInfo as OpenapiContentInfo}
                            contentOnly
                            maxWidth="calc(100vw - 346px)"
                            gitRef={gitRef}
                          />
                        </Render>
                      </Container>
                    </Render>
                  </Container>
                </Layout.Vertical>
              </Container>
            )
          },
          {
            id: FileSection.BLAME,
            title: getString('blame'),
            panel: (
              <Container className={css.gitBlame}>
                {[resourcePath + gitRef].map(key => (
                  <GitBlame repoMetadata={repoMetadata} resourcePath={resourcePath} gitRef={gitRef} key={key} />
                ))}
              </Container>
            )
          }
        ]}
      />
    </Container>
  )
}