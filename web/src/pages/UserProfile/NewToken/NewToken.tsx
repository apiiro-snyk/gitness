import React, { useMemo, useState } from 'react'
import {
  Button,
  ButtonVariation,
  Color,
  Container,
  Dialog,
  FlexExpander,
  FontVariation,
  FormikForm,
  FormInput,
  Layout,
  Text
} from '@harness/uicore'
import { useModalHook } from '@harness/use-modal'
import { Formik } from 'formik'
import { useMutate } from 'restful-react'
import moment from 'moment'
import * as Yup from 'yup'
import { Else, Match, Render, Truthy } from 'react-jsx-match'

import { useStrings } from 'framework/strings'
import { REGEX_VALID_REPO_NAME } from 'utils/Utils'
import { CodeIcon } from 'utils/GitUtils'
import { CopyButton } from 'components/CopyButton/CopyButton'
import { FormInputWithCopyButton } from 'components/UserManagementFlows/AddUserModal'

import css from 'components/CloneCredentialDialog/CloneCredentialDialog.module.scss'

const useNewToken = ({ onClose }: { onClose: () => void }) => {
  const { getString } = useStrings()
  const { mutate } = useMutate({ path: '/api/v1/user/tokens', verb: 'POST' })

  const [generatedToken, setGeneratedToken] = useState<string>()
  const isTokenGenerated = Boolean(generatedToken)

  const lifeTimeOptions = useMemo(
    () => [
      { label: getString('nDays', { number: 7 }), value: 604800000000000 },
      { label: getString('nDays', { number: 30 }), value: 2592000000000000 },
      { label: getString('nDays', { number: 60 }), value: 5184000000000000 },
      { label: getString('nDays', { number: 90 }), value: 7776000000000000 }
    ],
    [getString]
  )

  const onModalClose = () => {
    hideModal()
    onClose()
    setGeneratedToken()
  }

  const [openModal, hideModal] = useModalHook(() => {
    return (
      <Dialog isOpen enforceFocus={false} onClose={onModalClose} title={getString('createNewToken')}>
        <Formik
          initialValues={{
            uid: '',
            lifeTime: 0
          }}
          validationSchema={Yup.object().shape({
            uid: Yup.string()
              .required(getString('validation.nameIsRequired'))
              .matches(REGEX_VALID_REPO_NAME, getString('validation.nameInvalid')),
            lifeTime: Yup.number().required(getString('validation.expirationDateRequired'))
          })}
          onSubmit={async values => {
            const res = await mutate(values)
            setGeneratedToken(res?.access_token)
          }}>
          {formikProps => {
            const expiresAtString = moment(Date.now() + formikProps.values.lifeTime / 1000000).format(
              'dddd, MMMM DD YYYY'
            )

            return (
              <FormikForm>
                <FormInputWithCopyButton
                  name="uid"
                  label={getString('name')}
                  placeholder={getString('newToken.namePlaceholder')}
                  disabled={isTokenGenerated}
                />
                <FormInput.Select
                  name="lifeTime"
                  label={getString('expiration')}
                  items={lifeTimeOptions}
                  usePortal
                  disabled={isTokenGenerated}
                />
                {formikProps.values.lifeTime ? (
                  <Text
                    font={{ variation: FontVariation.SMALL_SEMI }}
                    color={Color.GREY_400}
                    margin={{ bottom: 'medium' }}>
                    {getString('newToken.expireOn', { date: expiresAtString })}
                  </Text>
                ) : null}
                <Render when={isTokenGenerated}>
                  <Text padding={{ bottom: 'small' }} font={{ variation: FontVariation.FORM_LABEL, size: 'small' }}>
                    {getString('token')}
                  </Text>
                  <Container padding={{ bottom: 'medium' }}>
                    <Layout.Horizontal className={css.layout}>
                      <Text className={css.url}>{generatedToken}</Text>
                      <FlexExpander />
                      <CopyButton
                        content={generatedToken || ''}
                        id={css.cloneCopyButton}
                        icon={CodeIcon.Copy}
                        iconProps={{ size: 14 }}
                      />
                    </Layout.Horizontal>
                  </Container>
                  <Text padding={{ bottom: 'medium' }} font={{ variation: FontVariation.BODY2_SEMI, size: 'small' }}>
                    {getString('newToken.tokenHelptext')}
                  </Text>
                </Render>
                <Match expr={isTokenGenerated}>
                  <Truthy>
                    <Button text={getString('close')} variation={ButtonVariation.TERTIARY} onClick={onModalClose} />
                  </Truthy>
                  <Else>
                    <Layout.Horizontal margin={{ top: 'xxxlarge' }} spacing="medium">
                      <Button
                        text={getString('newToken.generateToken')}
                        type="submit"
                        variation={ButtonVariation.PRIMARY}
                      />
                      <Button text={getString('cancel')} onClick={hideModal} variation={ButtonVariation.TERTIARY} />
                    </Layout.Horizontal>
                  </Else>
                </Match>
              </FormikForm>
            )
          }}
        </Formik>
      </Dialog>
    )
  }, [generatedToken])

  return {
    openModal,
    hideModal
  }
}

export default useNewToken
