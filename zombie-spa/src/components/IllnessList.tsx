import React, { FunctionComponent } from 'react';
import { Link } from 'react-router-dom';

import { IllnessData, Illness } from '../models/illnesses';

interface IllnessListProps {
  items: IllnessData[];
}

interface IllnessListItemProps {
  illness: Illness;
}

const IllnessListItem: FunctionComponent<IllnessListItemProps> = ({
  illness,
}) => {
  const hospitalLink = `/illnesses/${illness.id}/severity-level/`;
  return (
    <Link to={hospitalLink}>
      <li>
        <span>{illness.name}</span>
      </li>
    </Link>
  );
};

const IllnessList: FunctionComponent<IllnessListProps> = (props) => {
  return (
    <ul>
      <p>Select an illness:</p>
      {props.items.map((item) => (
        <IllnessListItem
          key={`illness-${item.illness.id}`}
          illness={item.illness}
        />
      ))}
    </ul>
  );
};

export default IllnessList;
