import React from 'react';
import { FlatList, StyleSheet, Text, View, Dimensions } from 'react-native';
import { Searchbar } from 'react-native-paper';
import { WARNA_ABU_ABU, WARNA_UTAMA } from '../../utils/constant';
import SwitchSelector from "react-native-switch-selector";
import { PesananAktif } from '../../components';

const Pesanan = () => {
    const [searchQuery, setSearchQuery] = React.useState('');
    const [orderList, setOrderList] = React.useState([
        { orderId: "0001212", status: "Sudah Selesai" },
        { orderId: "0001213", status: "Masih Dicuci" },
        { orderId: "0001214", status: "Sudah Selesai" },
        { orderId: "0001215", status: "Masih Dicuci" },
        { orderId: "0001216", status: "Sudah Selesai" },
        { orderId: "0001217", status: "Masih Dicuci" },
        { orderId: "0001218", status: "Sudah Selesai" },
        { orderId: "0001219", status: "Sudah Selesai" },
    ]);

    const options = [
        { label: "Aktif", value: "aktif", activeColor: WARNA_UTAMA },
        { label: "Selesai", value: "selesai", activeColor: WARNA_UTAMA },
        { label: "Semua", value: "semua", activeColor: WARNA_UTAMA }
    ];

    // fungsi untuk ngambil data search
    const onChangeSearch = (query) => {
        setSearchQuery(query)
        console.log(searchQuery)
    };

    // fungsi untuk switch data order
    const onShiftSwitch = (value) => {
        console.log(`Call onPress with value: ${value}`)
    }

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
                onPress={onShiftSwitch}
            />
            <View style={styles.container}>
                <FlatList
                    data={orderList}
                    renderItem={({ item }) => <PesananAktif orderId={item.orderId} status={item.status}/>}
                    showsVerticalScrollIndicator={false}
                />
            </View>
        </View>
    )
}

export default Pesanan

const windowWidth = Dimensions.get('window').width

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
    searchBar: {
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
    container: {
        flex: 1,
        paddingHorizontal: windowWidth * 0.035,
        marginTop: 10
    }
})
