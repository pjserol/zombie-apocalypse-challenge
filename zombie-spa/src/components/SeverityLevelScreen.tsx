import React, { FunctionComponent } from 'react';
import { useParams } from 'react-router-dom';
import SeverityLevelForm from './SeverityLevelForm';

const SeverityLevelScreen: FunctionComponent = () => {
  const { illnessId } = useParams();
  return (
    <React.Fragment>
      <SeverityLevelForm illnessId={illnessId} />
    </React.Fragment>
  );
};

export default SeverityLevelScreen;
