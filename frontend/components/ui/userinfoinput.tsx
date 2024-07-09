'use client';
import React, { useState, KeyboardEvent } from 'react';
import DOMPurify from 'dompurify';

const UserInfoInput: React.FC = () => {
  const [UserInfo, setUserInfo] = useState<string>('');
  const [UserInfosList, setUserInfosList] = useState<string[]>([]);

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUserInfo(event.target.value);
  };

  const handleAddUserInfo = () => {
    if (UserInfo.trim() !== '') {
      setUserInfosList([...UserInfosList, UserInfo.trim()]);
      setUserInfo('');
    }
  };

  const handleKeyPress = (event: KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      handleAddUserInfo();
    }
  };

  const handleDeleteUserInfo = (index: number) => {
    const updatedList = [...UserInfosList];
    updatedList.splice(index, 1);
    setUserInfosList(updatedList);
  };

  const renderUserInfos = () => {
    return UserInfosList.map((phone, index) => {
      // Regex to find URLs in text
      const urlRegex = /(https?:\/\/[^\s]+)/g;
      // Replace URLs with clickable links
      const userInfoWithLinks = phone.replace(urlRegex, (url) => `<a href="${url}" target="_blank" style="color: #0070f3; text-decoration: underline;">${url}</a>`);
      const sanitizedHtml = DOMPurify.sanitize(userInfoWithLinks);


      DOMPurify.addHook('afterSanitizeAttributes', function(node) {
        // Разрешаем только <a> теги с href атрибутом
        if (node.tagName.toLowerCase() === 'a') {
          const href = node.getAttribute('href');
          if (!/^https?:\/\//i.test(href)) {
            node.removeAttribute('href'); // Удаляем ссылки без HTTPS или HTTP
          }
        }
      });

      // Render HTML using dangerouslySetInnerHTML to handle the HTML tags
      return (
        <li key={index} style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', color: '090e20' }}>
          <div dangerouslySetInnerHTML={{ __html: sanitizedHtml }} />
          <button style={{ fontSize: '14px', color: '#555'}} onClick={() => handleDeleteUserInfo(index)}>Удалить</button>
        </li>
      );
    });
  };

  return (
    <div>
      <input
        type="text"
        value={UserInfo}
        onChange={handleInputChange}
        onKeyDown={handleKeyPress}
        placeholder="Раскажите о себе"
        style={{ marginRight: '10px' }}
      />
      <button style={{ fontSize: '14px', color: '#555'}} onClick={handleAddUserInfo}>Добавить</button>
      <ul>{renderUserInfos()}</ul>
    </div>
  );
};

export default UserInfoInput;
