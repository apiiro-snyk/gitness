import React from 'react'
import { Container, PageBody } from '@harness/uicore'
import { useGetRepositoryMetadata } from 'hooks/useGetRepositoryMetadata'
import { useGetResourceContent } from 'hooks/useGetResourceContent'
import { RepositoryFileEditHeader } from './RepositoryFileEditHeader/RepositoryFileEditHeader'
import { FileEditor } from './FileEditor/FileEditor'
import css from './RepositoryFileEdit.module.scss'

export default function RepositoryFileEdit() {
  const { gitRef, resourcePath, repoMetadata, error, loading, refetch } = useGetRepositoryMetadata()
  const {
    data: resourceContent,
    error: resourceError,
    loading: resourceLoading
  } = useGetResourceContent({ repoMetadata, gitRef, resourcePath })

  return (
    <Container className={css.main}>
      <PageBody loading={loading || resourceLoading} error={error || resourceError} retryOnError={() => refetch()}>
        {repoMetadata && resourceContent ? (
          <>
            <RepositoryFileEditHeader repoMetadata={repoMetadata} resourceContent={resourceContent} />
            <Container className={css.resourceContent}>
              <FileEditor
                repoMetadata={repoMetadata}
                gitRef={gitRef}
                resourcePath={resourcePath}
                resourceContent={resourceContent}
              />
            </Container>
          </>
        ) : null}
      </PageBody>
    </Container>
  )
}