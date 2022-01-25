import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import { Searchbar } from 'react-native-paper';
import { WARNA_ABU_ABU, WARNA_UTAMA } from '../../utils/constant';
import SwitchSelector from "react-native-switch-selector";

const Pesanan = () => {
    const [searchQuery, setSearchQuery] = React.useState('');

    const options = [
        { label: "Aktif", value: "1", activeColor: WARNA_UTAMA },
        { label: "Selesai", value: "1.5", activeColor: WARNA_UTAMA },
        { label: "Semua", value: "2", activeColor: WARNA_UTAMA }
      ];

    const onChangeSearch = (query) => {
        console.log(query)
        setSearchQuery(query)
    };

    return (
        <View style={styles.page}>
            <View style={styles.searchBox}>
                <Searchbar
                    style={styles.searchBar}
                    placeholder="Masukkan no. pesanan"
                    onChangeText={onChangeSearch}
                    value={searchQuery}
                    inputStyle={styles.placeholder}
                />
            </View>
            <SwitchSelector
                style={styles.selector}
                textStyle={styles.textUnselect}
                selectedTextStyle={styles.textSelected}
                options={options}
                initial={0}
                onPress={value => console.log(`Call onPress with value: ${value}`)}
            />
        </View>
    )
}

export default Pesanan

const styles = StyleSheet.create({
    page: {
        flex: 1,
        backgroundColor: WARNA_ABU_ABU,
    },
    searchBox: {
        backgroundColor: WARNA_UTAMA,
        paddingHorizontal: 20,
        paddingVertical: 10
    },
    searchBar:{
        borderRadius: 50,
        height: 40,
    },
    placeholder: {
        fontSize: 18,
        fontFamily: 'Roboto',
    },
    selector: {
        marginTop: 11,
        marginHorizontal: 12
    },
    textUnselect: {
        fontFamily: 'TitilliumWeb-SemiBold'
    },
    textSelected: {
        fontFamily: 'TitilliumWeb-SemiBold'
    },
})
