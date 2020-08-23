import React, { FunctionComponent } from 'react';

interface SeverityLevelFormProps {
  illnessId: number;
}

const SeverityLevelForm: FunctionComponent<SeverityLevelFormProps> = (
  props,
) => {
  console.log(props.illnessId);
  return <span>In construction...</span>;
};

export default SeverityLevelForm;
