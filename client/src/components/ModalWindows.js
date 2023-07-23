import React, { useEffect, useState } from "react";
import axios from "axios";
import styles from "./styles/modal.module.css";

const ModalWindow = ({ isOpen, onClose }) => {
    const [options, setOptions] = useState([]);
    const [selectedValue, setSelectedValue] = useState("");
    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get("http://localhost:5000/get-roles"); // замените на ваш URL сервера
                console.log(response.data);
                const flattenedOptions = response.data.flatMap((item) => {
                    return {
                        id: item.id,
                        name: item.name,
                        user_count: item.user_count,
                    };
                });
                setOptions(flattenedOptions);
            } catch (error) {
                console.error(error);
            }
        };
        fetchData();
    }, []);

    const handleSelectChange = (event) => {
        setSelectedValue(event.target.value);
    };

    const handleButtonClick = () => {
        console.log(selectedValue);
    };

    const changeRole = async () => {
        try {
            // Ваш запрос POST с выбранной ролью
            // const response = await axios.post("", { role: selectedRole });
            if (selectedValue.length === 0) {
                console.log("stop it")
                return 0
            }
            console.log("Role has been changed: " + selectedValue);
        } catch (e) {
            console.log(e);
        }
    };
    return (
        <>
            {isOpen && (
                <div className={styles.modal}>
                    <div className={styles.modalContent}>
                        <option disabled>Выберите роль</option>
                        <select onChange={handleSelectChange}>
                            <option value="">--Выберите роль--</option>
                            {options.map((option) => (
                                <option key={option.id} value={option.id}>
                                    {option.name}
                                </option>
                            ))}
                        </select>
                        <button onClick={handleButtonClick}>Вывести в консоль</button>
                        <button onClick={changeRole}>Сменить роль</button>
                        <button className={styles.closeButton} onClick={onClose}>Закрыть</button>
                    </div>
                </div>
            )}
        </>
    );
};

export default ModalWindow;